package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/layer-3/nitewatch/chain"
)

func main() {
	// 1. Setup Accounts
	// We generate keys for Admin, NeoDAX, Nitewatch, and a User.
	keys := make([]*ecdsa.PrivateKey, 4)
	addrs := make([]common.Address, 4)
	auths := make([]*bind.TransactOpts, 4)
	
	alloc := make(types.GenesisAlloc)
	balance := new(big.Int).Mul(big.NewInt(1000), big.NewInt(1e18)) // 1000 ETH

	fmt.Println("Setting up accounts...")
	for i := 0; i < 4; i++ {
		key, err := crypto.GenerateKey()
		if err != nil {
			log.Fatalf("Failed to generate key: %v", err)
		}
		keys[i] = key
		addrs[i] = crypto.PubkeyToAddress(key.PublicKey)
		
		auth, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(1337))
		if err != nil {
			log.Fatalf("Failed to create auth: %v", err)
		}
		auths[i] = auth

		alloc[addrs[i]] = types.Account{Balance: balance}
	}

	adminAuth := auths[0]
	neodaxAuth := auths[1]
	nitewatchAuth := auths[2]
	userAuth := auths[3]

	// 2. Setup Simulated Backend
	fmt.Println("Initializing simulated backend...")
	sim := backends.NewSimulatedBackend(alloc, 8000000)
	defer sim.Close()

	// 3. Deploy Contract
	fmt.Println("Deploying SimpleCustody contract...")
	custodyAddr, _, custody, err := chain.DeploySimpleCustody(adminAuth, sim, addrs[0], addrs[1], addrs[2])
	if err != nil {
		log.Fatalf("Failed to deploy SimpleCustody: %v", err)
	}
	sim.Commit()
	fmt.Printf("Contract deployed at: %s\n", custodyAddr.Hex())

	// 4. Start Event Listeners (Background)
	go func() {
		// Watch for Deposits
		depositCh := make(chan *chain.SimpleCustodyDeposited)
		depositSub, err := custody.WatchDeposited(&bind.WatchOpts{}, depositCh, nil, nil)
		if err != nil {
			log.Printf("Failed to watch deposits: %v", err)
			return
		}
		defer depositSub.Unsubscribe()

		// Watch for Withdrawal Requests
		withdrawStartedCh := make(chan *chain.SimpleCustodyWithdrawStarted)
		withdrawStartedSub, err := custody.WatchWithdrawStarted(&bind.WatchOpts{}, withdrawStartedCh, nil, nil, nil)
		if err != nil {
			log.Printf("Failed to watch withdraw started: %v", err)
			return
		}
		defer withdrawStartedSub.Unsubscribe()

		// Watch for Finalized Withdrawals
		withdrawFinalizedCh := make(chan *chain.SimpleCustodyWithdrawFinalized)
		withdrawFinalizedSub, err := custody.WatchWithdrawFinalized(&bind.WatchOpts{}, withdrawFinalizedCh, nil)
		if err != nil {
			log.Printf("Failed to watch withdraw finalized: %v", err)
			return
		}
		defer withdrawFinalizedSub.Unsubscribe()

		fmt.Println("Listening for events...")

		for {
			select {
			case ev := <-depositCh:
				fmt.Printf("[EVENT] Deposited: User=%s Token=%s Amount=%s\n", ev.User.Hex(), ev.Token.Hex(), ev.Amount.String())
			case ev := <-withdrawStartedCh:
				fmt.Printf("[EVENT] WithdrawStarted: ID=%x User=%s Amount=%s\n", ev.WithdrawalId, ev.User.Hex(), ev.Amount.String())
			case ev := <-withdrawFinalizedCh:
				fmt.Printf("[EVENT] WithdrawFinalized: ID=%x Success=%v\n", ev.WithdrawalId, ev.Success)
			case err := <-depositSub.Err():
				log.Printf("Deposit subscription error: %v", err)
				return
			case err := <-withdrawStartedSub.Err():
				log.Printf("WithdrawStarted subscription error: %v", err)
				return
			case err := <-withdrawFinalizedSub.Err():
				log.Printf("WithdrawFinalized subscription error: %v", err)
				return
			}
		}
	}()

	// Allow listeners to subscribe
	time.Sleep(100 * time.Millisecond)

	// 5. Execute Deposit
	fmt.Println("\n--- Executing Deposit ---")
	depositAmount := big.NewInt(1e18) // 1 ETH
	userAuth.Value = depositAmount
	tx, err := custody.Deposit(userAuth, common.Address{}, depositAmount)
	userAuth.Value = nil // Reset value
	if err != nil {
		log.Fatalf("Failed to deposit: %v", err)
	}
	fmt.Printf("Deposit transaction sent: %s\n", tx.Hash().Hex())
	sim.Commit()
	
	// Wait a bit for event processing
	time.Sleep(200 * time.Millisecond)

	// 6. Execute Withdrawal
	fmt.Println("\n--- Executing Withdrawal ---")
	withdrawAmount := big.NewInt(5e17) // 0.5 ETH
	nonce := big.NewInt(1)

	// NeoDAX starts withdrawal
	fmt.Println("NeoDAX initiating withdrawal...")
	tx, err = custody.StartWithdraw(neodaxAuth, addrs[3], common.Address{}, withdrawAmount, nonce)
	if err != nil {
		log.Fatalf("Failed to start withdraw: %v", err)
	}
	sim.Commit()
	
	receipt, err := sim.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		log.Fatalf("Failed to get receipt: %v", err)
	}

	// Find the withdrawal ID from the receipt logs to pass to finalize
	// In a real app, the Nitewatch daemon would pick this up from the event stream.
	// Here we parse it manually for the "Nitewatch" actor simulation.
	var withdrawalId [32]byte
	for _, log := range receipt.Logs {
		event, err := custody.ParseWithdrawStarted(*log)
		if err == nil {
			withdrawalId = event.WithdrawalId
			break
		}
	}
	fmt.Printf("Withdrawal ID: %x\n", withdrawalId)

	// Wait for event listener to print "WithdrawStarted"
	time.Sleep(200 * time.Millisecond)

	// Nitewatch finalizes withdrawal
	fmt.Println("Nitewatch finalizing withdrawal...")
	tx, err = custody.FinalizeWithdraw(nitewatchAuth, withdrawalId)
	if err != nil {
		log.Fatalf("Failed to finalize withdraw: %v", err)
	}
	sim.Commit()

	// Wait for event listener to print "WithdrawFinalized"
	time.Sleep(200 * time.Millisecond)

	fmt.Println("\n--- Demo Complete ---")
}
