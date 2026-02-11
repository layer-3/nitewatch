package main

import (
	"context"
	"crypto/ecdsa"
	"flag"
	"fmt"
	"log"
	"math/big"
	"net"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/layer-3/nitewatch/chain"
	pb "github.com/layer-3/nitewatch/proto"
)

type server struct {
	pb.UnimplementedCageServiceServer
	client  *ethclient.Client
	custody *chain.ICustody
	chainID *big.Int
	privKey *ecdsa.PrivateKey
}

func (s *server) StartWithdraw(ctx context.Context, req *pb.StartWithdrawRequest) (*pb.StartWithdrawResponse, error) {
	// 1. Parse and Validate Inputs
	if !common.IsHexAddress(req.UserAddress) {
		return nil, fmt.Errorf("invalid user address: %s", req.UserAddress)
	}
	user := common.HexToAddress(req.UserAddress)

	if !common.IsHexAddress(req.TokenAddress) {
		return nil, fmt.Errorf("invalid token address: %s", req.TokenAddress)
	}
	token := common.HexToAddress(req.TokenAddress)

	amount, ok := new(big.Int).SetString(req.Amount, 10)
	if !ok {
		return nil, fmt.Errorf("invalid amount: %s", req.Amount)
	}

	nonce, ok := new(big.Int).SetString(req.Nonce, 10)
	if !ok {
		return nil, fmt.Errorf("invalid nonce: %s", req.Nonce)
	}

	log.Printf("Received StartWithdraw request: User=%s Token=%s Amount=%s Nonce=%s",
		user.Hex(), token.Hex(), amount.String(), nonce.String())

	// 2. Prepare Transaction Auth
	auth, err := bind.NewKeyedTransactorWithChainID(s.privKey, s.chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create transactor: %v", err)
	}
	auth.Context = ctx

	// 3. Send Transaction
	tx, err := s.custody.StartWithdraw(auth, user, token, amount, nonce)
	if err != nil {
		return nil, fmt.Errorf("failed to submit startWithdraw transaction: %v", err)
	}

	log.Printf("Transaction sent: %s", tx.Hash().Hex())

	// 4. Wait for Mining to get WithdrawalID
	receipt, err := bind.WaitMined(ctx, s.client, tx)
	if err != nil {
		return nil, fmt.Errorf("failed to wait for transaction mining: %v", err)
	}

	if receipt.Status == 0 {
		return nil, fmt.Errorf("transaction reverted: %s", tx.Hash().Hex())
	}

	// 5. Extract WithdrawalID from Logs
	var withdrawalId [32]byte
	found := false

	for _, logMsg := range receipt.Logs {
		event, err := s.custody.ParseWithdrawStarted(*logMsg)
		if err == nil {
			withdrawalId = event.WithdrawalId
			found = true
			break
		}
	}

	if !found {
		return nil, fmt.Errorf("WithdrawStarted event not found in transaction logs")
	}

	log.Printf("Withdrawal started successfully. ID: %x", withdrawalId)

	return &pb.StartWithdrawResponse{
		WithdrawalId: common.Hash(withdrawalId).Hex(),
		TxHash:       tx.Hash().Hex(),
	}, nil
}

func main() {
	rpcURL := flag.String("rpc", "ws://127.0.0.1:8545", "Ethereum RPC URL")
	contractAddr := flag.String("contract", "", "ICustody contract address")
	privateKeyHex := flag.String("key", "", "Private key for Cage (hex)")
	port := flag.Int("port", 50051, "gRPC server port")
	flag.Parse()

	if *contractAddr == "" || *privateKeyHex == "" {
		log.Fatal("Contract address and private key are required")
	}

	// Parse private key once at startup
	key, err := crypto.HexToECDSA(*privateKeyHex)
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}

	// Initialize Ethereum Client
	client, err := ethclient.Dial(*rpcURL)
	if err != nil {
		log.Fatalf("Failed to connect to RPC: %v", err)
	}
	defer client.Close()

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatalf("Failed to get chain ID: %v", err)
	}

	// Initialize Contract Binding
	custody, err := chain.NewICustody(common.HexToAddress(*contractAddr), client)
	if err != nil {
		log.Fatalf("Failed to bind contract: %v", err)
	}

	// Start gRPC Server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCageServiceServer(s, &server{
		client:  client,
		custody: custody,
		chainID: chainID,
		privKey: key,
	})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	log.Printf("Cage service listening on :%d", *port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
