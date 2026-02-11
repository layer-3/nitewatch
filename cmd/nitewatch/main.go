package main

import (
	"context"
	_ "embed"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	nw "github.com/layer-3/nitewatch"
	"github.com/layer-3/nitewatch/chain"
	"github.com/layer-3/nitewatch/core"
	"github.com/layer-3/nitewatch/store"
)

//go:embed limits.yaml
var limitsConfig []byte

func main() {
	rpcURL := flag.String("rpc", "ws://127.0.0.1:8545", "Ethereum RPC URL (WebSocket required)")
	contractAddr := flag.String("contract", "", "ICustody contract address")
	privateKeyHex := flag.String("key", "", "Private key for finalizing withdrawals (hex)")
	confirmations := flag.Uint64("confirmations", 12, "Number of block confirmations to wait")
	dbPath := flag.String("db", "nitewatch.db", "Path to SQLite database")
	flag.Parse()

	if *contractAddr == "" || *privateKeyHex == "" {
		log.Fatal("Contract address and private key are required")
	}

	// 1. Initialize Store
	gormDB, err := gorm.Open(sqlite.Open(*dbPath), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	db, err := store.NewAdapter(gormDB)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// 2. Load Configuration and Initialize Checker
	var cfg core.Config
	if err := yaml.Unmarshal(limitsConfig, &cfg); err != nil {
		log.Fatalf("Failed to parse embedded limits.yaml: %v", err)
	}

	checker, err := core.NewChecker(cfg, db)
	if err != nil {
		log.Fatalf("Failed to initialize checker: %v", err)
	}

	// 3. Connect to Ethereum
	client, err := ethclient.Dial(*rpcURL)
	if err != nil {
		log.Fatalf("Failed to connect to RPC: %v", err)
	}
	defer client.Close()

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatalf("Failed to get chain ID: %v", err)
	}

	// 4. Setup Signer
	key, err := crypto.HexToECDSA(*privateKeyHex)
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		log.Fatalf("Failed to create transactor: %v", err)
	}

	// 5. Setup Contract Binding
	addr := common.HexToAddress(*contractAddr)
	custodyContract, err := chain.NewICustody(addr, client)
	if err != nil {
		log.Fatalf("Failed to bind contract: %v", err)
	}

	// 6. Setup Listener
	listener := chain.NewListener(client, custodyContract, *confirmations)

	// 7. Start Watching
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	withdrawals := make(chan *chain.ICustodyWithdrawStarted)

	// Handle shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigCh
		log.Println("Shutting down...")
		cancel()
	}()

	go func() {
		log.Println("Listening for WithdrawStarted events...")
		if err := listener.WatchWithdrawStarted(ctx, withdrawals); err != nil {
			if ctx.Err() == nil {
				log.Printf("WatchWithdrawStarted error: %v", err)
			}
		}
	}()

	// 8. Process Loop
	for event := range withdrawals {
		log.Printf("New withdrawal request: ID=%x User=%s Token=%s Amount=%s",
			event.WithdrawalId, event.User.Hex(), event.Token.Hex(), event.Amount)

		// Check Limits
		if err := checker.Check(event.Token, event.Amount); err != nil {
			log.Printf("Withdrawal %x blocked by policy: %v", event.WithdrawalId, err)
			continue
		}

		// Finalize
		txAuth := *auth
		txAuth.Context = ctx

		tx, err := custodyContract.FinalizeWithdraw(&txAuth, event.WithdrawalId)
		if err != nil {
			log.Printf("Failed to finalize withdrawal %x: %v", event.WithdrawalId, err)
			continue
		}

		log.Printf("Sent finalize tx: %s for withdrawal %x", tx.Hash().Hex(), event.WithdrawalId)

		receipt, err := bind.WaitMined(ctx, client, tx)
		if err != nil {
			log.Printf("Transaction mining failed: %v", err)
			continue
		}

		if receipt.Status == 1 {
			log.Printf("Withdrawal %x finalized successfully on-chain.", event.WithdrawalId)

			// Record usage in DB
			record := &nw.Withdrawal{
				WithdrawalID: event.WithdrawalId,
				User:         event.User,
				Token:        event.Token,
				Amount:       event.Amount,
				BlockNumber:  receipt.BlockNumber.Uint64(),
				TxHash:       tx.Hash(),
				Timestamp:    time.Now(),
			}

			if err := checker.Record(record); err != nil {
				log.Printf("Failed to record withdrawal %x in DB: %v", event.WithdrawalId, err)
			}
		} else {
			log.Printf("Withdrawal %x finalization tx failed (reverted).", event.WithdrawalId)
		}
	}
}
