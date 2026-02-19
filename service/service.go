package service

import (
	"context"
	"fmt"
	"log/slog"
	"math/big"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/layer-3/nitewatch/config"
	"github.com/layer-3/nitewatch/custody"
	"github.com/layer-3/nitewatch/internal/checker"
	"github.com/layer-3/nitewatch/internal/store"
)

type httpServer struct {
	Engine *gin.Engine
	server *http.Server
}

func newHTTPServer(addr string) *httpServer {
	engine := gin.New()
	engine.Use(gin.Recovery())
	return &httpServer{
		Engine: engine,
		server: &http.Server{Addr: addr, Handler: engine},
	}
}

func (s *httpServer) Run() error                         { return s.server.ListenAndServe() }
func (s *httpServer) Shutdown(ctx context.Context) error { return s.server.Shutdown(ctx) }

type Service struct {
	Config config.Config
	Logger *slog.Logger

	web       *httpServer
	ethClient custody.EthBackend
	contract  *custody.IWithdraw
	listener  *custody.Listener
	auth      *bind.TransactOpts
	checker   *checker.Checker
	store     *store.Adapter

	workerReady int32
}

// New creates a Service that dials an Ethereum node via the configured RPC URL.
func New(conf config.Config) (*Service, error) {
	client, err := ethclient.Dial(conf.Blockchain.RPCURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum RPC: %w", err)
	}

	svc, err := NewWithBackend(conf, client)
	if err != nil {
		client.Close()
		return nil, err
	}
	return svc, nil
}

// NewWithBackend creates a Service using a pre-existing Ethereum backend.
// The caller is responsible for closing the backend when done.
func NewWithBackend(conf config.Config, client custody.EthBackend) (*Service, error) {
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil)).With("service", "nitewatch")

	srv := newHTTPServer(conf.ListenAddr)

	gormDB, err := gorm.Open(sqlite.Open(conf.DBPath), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	db, err := store.NewAdapter(gormDB)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}

	globalLimits, err := parseLimitsConfig(conf.Limits)
	if err != nil {
		return nil, fmt.Errorf("failed to parse global limits: %w", err)
	}

	userOverrides, err := parseUserOverrides(conf.PerUserOverrides)
	if err != nil {
		return nil, fmt.Errorf("failed to parse per-user overrides: %w", err)
	}

	chk := checker.New(globalLimits, userOverrides, db)

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID: %w", err)
	}

	key, err := crypto.HexToECDSA(conf.Blockchain.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create transactor: %w", err)
	}

	addr := common.HexToAddress(conf.Blockchain.ContractAddr)
	withdrawContract, err := custody.NewIWithdraw(addr, client)
	if err != nil {
		return nil, fmt.Errorf("failed to bind IWithdraw contract: %w", err)
	}

	listener := custody.NewListener(client, addr, withdrawContract, nil)

	return &Service{
		Config:    conf,
		Logger:    logger,
		web:       srv,
		ethClient: client,
		contract:  withdrawContract,
		listener:  listener,
		auth:      auth,
		checker:   chk,
		store:     db,
	}, nil
}

func (svc *Service) IsWorkerReady() bool {
	return atomic.LoadInt32(&svc.workerReady) == 1
}

func (svc *Service) setWorkerReady() {
	atomic.StoreInt32(&svc.workerReady, 1)
}

func (svc *Service) RunWorker() error {
	return svc.RunWorkerWithContext(context.Background())
}

func (svc *Service) RunWorkerWithContext(ctx context.Context) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		svc.Logger.Info("Starting health endpoint server")
		return svc.web.Run()
	})

	g.Go(func() error {
		fromBlock, fromLogIdx, err := svc.store.GetCursor("withdraw_started")
		if err != nil {
			svc.Logger.Warn("Failed to read withdraw_started cursor, starting from head", "error", err)
		}
		svc.Logger.Info("Starting WithdrawStarted event watcher", "from_block", fromBlock, "from_log_index", fromLogIdx)
		withdrawals := make(chan *custody.WithdrawStartedEvent)
		go svc.listener.WatchWithdrawStarted(ctx, withdrawals, fromBlock, fromLogIdx)
		for event := range withdrawals {
			svc.processWithdrawal(ctx, event)
		}
		return nil
	})

	g.Go(func() error {
		<-ctx.Done()
		svc.Logger.Info("Shutting down health endpoint server")
		ctxShutdown, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelShutdown()
		return svc.web.Shutdown(ctxShutdown)
	})

	g.Go(func() error {
		<-ctx.Done()
		svc.Logger.Info("Closing Ethereum client")
		svc.ethClient.Close()
		return nil
	})

	svc.setWorkerReady()
	svc.Logger.Info("Worker started")

	return g.Wait()
}

func (svc *Service) processWithdrawal(ctx context.Context, event *custody.WithdrawStartedEvent) {
	wID := common.Hash(event.WithdrawalID).Hex()
	logger := svc.Logger.With(
		"withdrawal_id", wID,
		"user", event.User.Hex(),
		"token", event.Token.Hex(),
		"amount", event.Amount,
	)

	if svc.store.HasWithdrawEvent(wID) {
		logger.Info("Event already processed, skipping")
		return
	}

	logger.Info("Processing withdrawal request")

	baseModel := store.WithdrawEventModel{
		WithdrawalID: wID,
		UserAddress:  event.User.Hex(),
		TokenAddress: event.Token.Hex(),
		Amount:       event.Amount.String(),
		BlockNumber:  event.BlockNumber,
		TxHash:       event.TxHash.Hex(),
		LogIndex:     uint(event.LogIndex),
	}

	if err := svc.checker.Check(event.User, event.Token, event.Amount); err != nil {
		logger.Warn("Withdrawal blocked by policy, rejecting", "reason", err)

		txAuth := *svc.auth
		txAuth.Context = ctx
		tx, txErr := svc.contract.RejectWithdraw(&txAuth, event.WithdrawalID)
		if txErr != nil {
			logger.Error("Failed to reject withdrawal", "error", txErr)
			baseModel.Decision = "error"
			baseModel.Reason = fmt.Sprintf("reject tx failed: %v", txErr)
			svc.recordEvent(logger, &baseModel)
			return
		}
		logger.Info("Sent reject transaction", "tx_hash", tx.Hash().Hex())
		if _, txErr = bind.WaitMined(ctx, svc.ethClient, tx); txErr != nil {
			logger.Error("Failed waiting for reject tx to be mined", "error", txErr)
			baseModel.Decision = "error"
			baseModel.Reason = fmt.Sprintf("reject tx mining failed: %v", txErr)
			svc.recordEvent(logger, &baseModel)
			return
		}

		baseModel.Decision = "rejected"
		baseModel.Reason = err.Error()
		svc.recordEvent(logger, &baseModel)
		return
	}

	txAuth := *svc.auth
	txAuth.Context = ctx

	tx, err := svc.contract.FinalizeWithdraw(&txAuth, event.WithdrawalID)
	if err != nil {
		logger.Error("Failed to finalize withdrawal", "error", err)
		baseModel.Decision = "error"
		baseModel.Reason = fmt.Sprintf("finalize tx failed: %v", err)
		svc.recordEvent(logger, &baseModel)
		return
	}

	logger.Info("Sent finalize transaction", "tx_hash", tx.Hash().Hex())

	receipt, err := bind.WaitMined(ctx, svc.ethClient, tx)
	if err != nil {
		logger.Error("Transaction mining failed", "error", err)
		baseModel.Decision = "error"
		baseModel.Reason = fmt.Sprintf("finalize tx mining failed: %v", err)
		svc.recordEvent(logger, &baseModel)
		return
	}

	if receipt.Status == 1 {
		logger.Info("Withdrawal finalized successfully on-chain")

		record := &custody.Withdrawal{
			WithdrawalID: event.WithdrawalID,
			User:         event.User,
			Token:        event.Token,
			Amount:       event.Amount,
			BlockNumber:  receipt.BlockNumber.Uint64(),
			TxHash:       tx.Hash(),
			Timestamp:    time.Now(),
		}
		if err := svc.checker.Record(record); err != nil {
			logger.Error("Failed to record withdrawal in DB", "error", err)
		}

		baseModel.Decision = "approved"
		svc.recordEvent(logger, &baseModel)
	} else {
		logger.Error("Withdrawal finalization tx reverted")
		baseModel.Decision = "error"
		baseModel.Reason = "finalize tx reverted on-chain"
		svc.recordEvent(logger, &baseModel)
	}
}

func (svc *Service) recordEvent(logger *slog.Logger, ev *store.WithdrawEventModel) {
	if err := svc.store.RecordWithdrawEvent(ev); err != nil {
		logger.Error("Failed to record withdraw event", "error", err)
	}
}

func parseLimitsConfig(lc config.LimitsConfig) (map[common.Address]checker.Limit, error) {
	limits := make(map[common.Address]checker.Limit)
	for addrStr, conf := range lc {
		if !common.IsHexAddress(addrStr) {
			return nil, fmt.Errorf("invalid address: %s", addrStr)
		}
		addr := common.HexToAddress(addrStr)

		l := checker.Limit{}
		if conf.Hourly != "" {
			val, ok := new(big.Int).SetString(conf.Hourly, 10)
			if !ok {
				return nil, fmt.Errorf("invalid hourly limit for %s: %s", addrStr, conf.Hourly)
			}
			l.Hourly = val
		}
		if conf.Daily != "" {
			val, ok := new(big.Int).SetString(conf.Daily, 10)
			if !ok {
				return nil, fmt.Errorf("invalid daily limit for %s: %s", addrStr, conf.Daily)
			}
			l.Daily = val
		}
		limits[addr] = l
	}
	return limits, nil
}

func parseUserOverrides(overrides map[string]config.LimitsConfig) (map[common.Address]map[common.Address]checker.Limit, error) {
	result := make(map[common.Address]map[common.Address]checker.Limit)
	for userAddrStr, tokenLimits := range overrides {
		if !common.IsHexAddress(userAddrStr) {
			return nil, fmt.Errorf("invalid user address in per_user_overrides: %s", userAddrStr)
		}
		userAddr := common.HexToAddress(userAddrStr)
		parsed, err := parseLimitsConfig(tokenLimits)
		if err != nil {
			return nil, fmt.Errorf("per-user overrides for %s: %w", userAddrStr, err)
		}
		result[userAddr] = parsed
	}
	return result, nil
}
