package chain

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func TestSimpleCustodyFlow(t *testing.T) {
	// 1. Setup Accounts
	// We need 4 accounts: Admin, NeoDAX, Nitewatch, User
	keys := make([]*ecdsa.PrivateKey, 4)
	addrs := make([]common.Address, 4)
	auths := make([]*bind.TransactOpts, 4)
	
	alloc := make(types.GenesisAlloc)
	balance := new(big.Int).Mul(big.NewInt(1000), big.NewInt(1e18)) // 1000 ETH

	for i := 0; i < 4; i++ {
		key, err := crypto.GenerateKey()
		if err != nil {
			t.Fatalf("Failed to generate key: %v", err)
		}
		keys[i] = key
		addrs[i] = crypto.PubkeyToAddress(key.PublicKey)
		
		auth, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(1337))
		if err != nil {
			t.Fatalf("Failed to create auth: %v", err)
		}
		auths[i] = auth

		alloc[addrs[i]] = types.Account{Balance: balance}
	}

	adminAuth := auths[0]
	neodaxAuth := auths[1]
	nitewatchAuth := auths[2]
	userAuth := auths[3]

	adminAddr := addrs[0]
	neodaxAddr := addrs[1]
	nitewatchAddr := addrs[2]
	userAddr := addrs[3]

	// 2. Setup Simulated Backend
	sim := backends.NewSimulatedBackend(alloc, 8000000)
	defer sim.Close()

	// 3. Deploy Contract
	custodyAddr, _, custody, err := DeploySimpleCustody(adminAuth, sim, adminAddr, neodaxAddr, nitewatchAddr)
	if err != nil {
		t.Fatalf("Failed to deploy SimpleCustody: %v", err)
	}
	sim.Commit()

	// 4. Test Deposit Flow
	t.Run("Deposit ETH", func(t *testing.T) {
		depositAmount := big.NewInt(1e18) // 1 ETH
		
		// User deposits
		// Note: Deposit requires value to be sent. `userAuth` needs to be updated with value.
		userAuth.Value = depositAmount
		tx, err := custody.Deposit(userAuth, common.Address{}, depositAmount)
		userAuth.Value = nil // Reset
		if err != nil {
			t.Fatalf("Failed to deposit: %v", err)
		}
		sim.Commit()

		receipt, err := sim.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			t.Fatalf("Failed to get receipt: %v", err)
		}
		if receipt.Status != types.ReceiptStatusSuccessful {
			t.Fatal("Deposit transaction failed")
		}

		// Check contract balance
		contractBalance, err := sim.BalanceAt(context.Background(), custodyAddr, nil)
		if err != nil {
			t.Fatalf("Failed to get contract balance: %v", err)
		}
		if contractBalance.Cmp(depositAmount) != 0 {
			t.Errorf("Expected contract balance %v, got %v", depositAmount, contractBalance)
		}
	})

	// 5. Test Withdrawal Flow
	t.Run("Withdraw ETH", func(t *testing.T) {
		withdrawAmount := big.NewInt(5e17) // 0.5 ETH
		nonce := big.NewInt(1)

		// Record User balance before
		userBalBefore, err := sim.BalanceAt(context.Background(), userAddr, nil)
		if err != nil {
			t.Fatalf("Failed to get user balance: %v", err)
		}

		// A. Start Withdraw (NeoDAX)
		tx, err := custody.StartWithdraw(neodaxAuth, userAddr, common.Address{}, withdrawAmount, nonce)
		if err != nil {
			t.Fatalf("Failed to start withdraw: %v", err)
		}
		sim.Commit()

		receipt, err := sim.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			t.Fatalf("Failed to get start withdraw receipt: %v", err)
		}
		if receipt.Status != types.ReceiptStatusSuccessful {
			t.Fatal("StartWithdraw transaction failed")
		}

		// Parse event to get withdrawalId
		// In a real scenario we'd parse logs. Here we can iterate logs.
		var withdrawalId [32]byte
		found := false
		for _, log := range receipt.Logs {
			event, err := custody.ParseWithdrawStarted(*log)
			if err == nil {
				withdrawalId = event.WithdrawalId
				found = true
				break
			}
		}
		if !found {
			t.Fatal("WithdrawStarted event not found")
		}

		// B. Finalize Withdraw (Nitewatch)
		tx, err = custody.FinalizeWithdraw(nitewatchAuth, withdrawalId)
		if err != nil {
			t.Fatalf("Failed to finalize withdraw: %v", err)
		}
		sim.Commit()

		receipt, err = sim.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			t.Fatalf("Failed to get finalize withdraw receipt: %v", err)
		}
		if receipt.Status != types.ReceiptStatusSuccessful {
			t.Fatal("FinalizeWithdraw transaction failed")
		}

		// Check User balance after
		userBalAfter, err := sim.BalanceAt(context.Background(), userAddr, nil)
		if err != nil {
			t.Fatalf("Failed to get user balance: %v", err)
		}

		expectedBal := new(big.Int).Add(userBalBefore, withdrawAmount)
		// Note: user pays gas for nothing here? No, user doesn't call anything.
		// User receives funds. Gas is paid by Nitewatch.
		// So user balance should be exactly previous + withdrawn amount.
		
		if userBalAfter.Cmp(expectedBal) != 0 {
			t.Errorf("Expected user balance %v, got %v", expectedBal, userBalAfter)
		}
	})
}
