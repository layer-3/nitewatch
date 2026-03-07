//go:build !short

package service_test

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"path/filepath"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/layer-3/nitewatch/config"
	"github.com/layer-3/nitewatch/custody"
	"github.com/layer-3/nitewatch/service"
)

const (
	nativeToken = "0x0000000000000000000000000000000000000000"
	// Simulated backend always uses chainID 1337.
	simChainID = 1337
)

type testEnv struct {
	sim      *simulated.Backend
	client   custody.EthBackend
	keys     [4]*ecdsa.PrivateKey
	addrs    [4]common.Address
	auths    [4]*bind.TransactOpts
	contract *custody.SimpleCustody
	addr     common.Address
}

func (e *testEnv) adminAuth() *bind.TransactOpts   { return e.auths[0] }
func (e *testEnv) neodaxAuth() *bind.TransactOpts  { return e.auths[1] }
func (e *testEnv) nitewatchKey() *ecdsa.PrivateKey { return e.keys[2] }
func (e *testEnv) nitewatchAddr() common.Address   { return e.addrs[2] }
func (e *testEnv) userAddr() common.Address        { return e.addrs[3] }

func newTestEnv(t *testing.T) *testEnv {
	t.Helper()

	keys := [4]*ecdsa.PrivateKey{}
	addrs := [4]common.Address{}
	auths := [4]*bind.TransactOpts{}
	alloc := make(types.GenesisAlloc)
	balance := new(big.Int).Mul(big.NewInt(1000), big.NewInt(1e18))

	for i := range 4 {
		key, err := crypto.GenerateKey()
		require.NoError(t, err)
		keys[i] = key
		addrs[i] = crypto.PubkeyToAddress(key.PublicKey)
		auth, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(simChainID))
		require.NoError(t, err)
		auths[i] = auth
		alloc[addrs[i]] = types.Account{Balance: balance}
	}

	sim := simulated.NewBackend(alloc)
	t.Cleanup(func() { sim.Close() })

	client := simBackendClient{Client: sim.Client(), backend: sim}

	contractAddr, tx, contract, err := custody.DeploySimpleCustody(
		auths[0], // admin deploys
		client,   // backend
		addrs[0], // admin
		addrs[1], // neodax
		addrs[2], // nitewatch
	)
	require.NoError(t, err)
	sim.Commit()

	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt.Status, "contract deployment failed")

	return &testEnv{
		sim:      sim,
		client:   client,
		keys:     keys,
		addrs:    addrs,
		auths:    auths,
		contract: contract,
		addr:     contractAddr,
	}
}

// simBackendClient wraps simulated.Client to add Close(), satisfying custody.EthBackend.
type simBackendClient struct {
	simulated.Client
	backend *simulated.Backend
}

func (c simBackendClient) Close() { c.backend.Close() }

// autoCommit mines blocks periodically so that bind.WaitMined can return.
func autoCommit(ctx context.Context, sim *simulated.Backend, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			sim.Commit()
		}
	}
}

func createNitewatchService(t *testing.T, env *testEnv, limitWei string) *service.Service {
	t.Helper()

	conf := config.Config{
		Blockchain: config.BlockchainConfig{
			ContractAddr:       env.addr.Hex(),
			PrivateKey:         fmt.Sprintf("%x", crypto.FromECDSA(env.nitewatchKey())),
			ConfirmationBlocks: 1,
			PollInterval:       200 * time.Millisecond,
		},
		Limits: config.LimitsConfig{
			nativeToken: config.LimitConfig{
				Hourly: limitWei,
				Daily:  limitWei,
			},
		},
		DBPath:     filepath.Join(t.TempDir(), "nitewatch.db"),
		ListenAddr: ":0",
	}

	svc, err := service.NewWithBackend(conf, env.client)
	require.NoError(t, err)
	return svc
}

func runNitewatchService(t *testing.T, svc *service.Service) {
	t.Helper()

	ctx, cancel := context.WithCancel(t.Context())
	errCh := make(chan error, 1)
	go func() { errCh <- svc.RunWorkerWithContext(ctx) }()

	require.Eventually(t, svc.IsWorkerReady, 30*time.Second, 100*time.Millisecond,
		"nitewatch worker did not become ready")

	t.Cleanup(func() {
		cancel()
		select {
		case err := <-errCh:
			if err != nil && !errors.Is(err, context.Canceled) {
				t.Logf("nitewatch worker stopped with error: %v", err)
			}
		case <-time.After(10 * time.Second):
			t.Log("nitewatch worker cleanup timeout")
		}
	})
}

// waitForWithdrawFinalized polls for a WithdrawFinalized event with the given success value.
func waitForWithdrawFinalized(t *testing.T, env *testEnv, timeout time.Duration) *custody.SimpleCustodyWithdrawFinalized {
	t.Helper()

	deadline := time.After(timeout)
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-deadline:
			t.Fatal("timed out waiting for WithdrawFinalized event")
			return nil
		case <-ticker.C:
			iter, err := env.contract.FilterWithdrawFinalized(&bind.FilterOpts{
				Start:   0,
				Context: context.Background(),
			}, nil)
			require.NoError(t, err)
			if iter.Next() {
				ev := iter.Event
				iter.Close()
				return ev
			}
			iter.Close()
		}
	}
}

func TestWithdrawalFinalized(t *testing.T) {
	env := newTestEnv(t)

	ctx, cancel := context.WithCancel(t.Context())
	defer cancel()
	go autoCommit(ctx, env.sim, 100*time.Millisecond)

	// 100 ETH limit — well above our 0.5 ETH withdrawal.
	svc := createNitewatchService(t, env, "100000000000000000000")
	runNitewatchService(t, svc)

	// User deposits 1 ETH.
	depositAmount := big.NewInt(1e18)
	userAuth := copyAuth(env.auths[3])
	userAuth.Value = depositAmount
	tx, err := env.contract.Deposit(userAuth, common.Address{}, depositAmount)
	require.NoError(t, err)
	env.sim.Commit()
	receipt, err := env.client.TransactionReceipt(context.Background(), tx.Hash())
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt.Status, "deposit tx failed")

	// Record user balance before withdrawal.
	userBalBefore, err := env.client.(ethereum.ChainStateReader).BalanceAt(
		context.Background(), env.userAddr(), nil)
	require.NoError(t, err)

	// NeoDAX starts a 0.5 ETH withdrawal for user.
	withdrawAmount := new(big.Int).Div(depositAmount, big.NewInt(2))
	neodaxAuth := copyAuth(env.neodaxAuth())
	tx, err = env.contract.StartWithdraw(neodaxAuth, env.userAddr(), common.Address{}, withdrawAmount, big.NewInt(1))
	require.NoError(t, err)
	env.sim.Commit()
	receipt, err = env.client.TransactionReceipt(context.Background(), tx.Hash())
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt.Status, "startWithdraw tx failed")

	// Wait for nitewatch to finalize the withdrawal.
	ev := waitForWithdrawFinalized(t, env, 30*time.Second)
	assert.True(t, ev.Success, "expected withdrawal to be finalized successfully")

	// Verify user received funds.
	userBalAfter, err := env.client.(ethereum.ChainStateReader).BalanceAt(
		context.Background(), env.userAddr(), nil)
	require.NoError(t, err)
	expected := new(big.Int).Add(userBalBefore, withdrawAmount)
	assert.Equal(t, expected.String(), userBalAfter.String(),
		"user balance should increase by the withdrawn amount")
}

func TestWithdrawalRejected(t *testing.T) {
	env := newTestEnv(t)

	ctx, cancel := context.WithCancel(t.Context())
	defer cancel()
	go autoCommit(ctx, env.sim, 100*time.Millisecond)

	// 0.1 ETH limit — below our 0.5 ETH withdrawal.
	svc := createNitewatchService(t, env, "100000000000000000")
	runNitewatchService(t, svc)

	// User deposits 1 ETH.
	depositAmount := big.NewInt(1e18)
	userAuth := copyAuth(env.auths[3])
	userAuth.Value = depositAmount
	tx, err := env.contract.Deposit(userAuth, common.Address{}, depositAmount)
	require.NoError(t, err)
	env.sim.Commit()
	receipt, err := env.client.TransactionReceipt(context.Background(), tx.Hash())
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt.Status, "deposit tx failed")

	// Record user balance before withdrawal.
	userBalBefore, err := env.client.(ethereum.ChainStateReader).BalanceAt(
		context.Background(), env.userAddr(), nil)
	require.NoError(t, err)

	// NeoDAX starts a 0.5 ETH withdrawal (exceeds 0.1 ETH limit).
	withdrawAmount := new(big.Int).Div(depositAmount, big.NewInt(2))
	neodaxAuth := copyAuth(env.neodaxAuth())
	tx, err = env.contract.StartWithdraw(neodaxAuth, env.userAddr(), common.Address{}, withdrawAmount, big.NewInt(1))
	require.NoError(t, err)
	env.sim.Commit()
	receipt, err = env.client.TransactionReceipt(context.Background(), tx.Hash())
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt.Status, "startWithdraw tx failed")

	// Wait for nitewatch to reject the withdrawal.
	ev := waitForWithdrawFinalized(t, env, 30*time.Second)
	assert.False(t, ev.Success, "expected withdrawal to be rejected")

	// Verify user balance unchanged.
	userBalAfter, err := env.client.(ethereum.ChainStateReader).BalanceAt(
		context.Background(), env.userAddr(), nil)
	require.NoError(t, err)
	assert.Equal(t, userBalBefore.String(), userBalAfter.String(),
		"user balance should not change after rejection")
}

// TestGasBufferApplied verifies that the gas buffer mechanism works correctly
// by deploying ThresholdCustody (the real contract where the bug occurs),
// comparing the bare eth_estimateGas result against the 20%-buffered value,
// and confirming the transaction succeeds with the buffered gas limit.
func TestGasBufferApplied(t *testing.T) {
	const gasBufferPercent = 20

	// --- Setup: 3 signers, threshold=2, deploy ThresholdCustody ---
	numAccounts := 4 // 3 signers + 1 user
	keys := make([]*ecdsa.PrivateKey, numAccounts)
	addrs := make([]common.Address, numAccounts)
	auths := make([]*bind.TransactOpts, numAccounts)
	alloc := make(types.GenesisAlloc)
	balance := new(big.Int).Mul(big.NewInt(1000), big.NewInt(1e18))

	for i := range numAccounts {
		key, err := crypto.GenerateKey()
		require.NoError(t, err)
		keys[i] = key
		addrs[i] = crypto.PubkeyToAddress(key.PublicKey)
		auth, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(simChainID))
		require.NoError(t, err)
		auths[i] = auth
		alloc[addrs[i]] = types.Account{Balance: balance}
	}

	sim := simulated.NewBackend(alloc)
	t.Cleanup(func() { sim.Close() })
	client := simBackendClient{Client: sim.Client(), backend: sim}

	signers := []common.Address{addrs[0], addrs[1], addrs[2]}
	contractAddr, tx, tcContract, err := custody.DeployThresholdCustody(
		copyAuth(auths[0]), client, signers, 2,
	)
	require.NoError(t, err)
	sim.Commit()
	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt.Status, "ThresholdCustody deployment failed")

	// Also bind as IWithdraw for FinalizeWithdraw.
	iwContract, err := custody.NewIWithdraw(contractAddr, client)
	require.NoError(t, err)

	// --- User deposits 1 ETH ---
	depositAmount := big.NewInt(1e18)
	userAuth := copyAuth(auths[3])
	userAuth.Value = depositAmount
	tx, err = tcContract.Deposit(userAuth, common.Address{}, depositAmount)
	require.NoError(t, err)
	sim.Commit()
	receipt, err = client.TransactionReceipt(context.Background(), tx.Hash())
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt.Status, "deposit failed")

	// --- Signer0 starts a withdrawal (counts as 1 approval) ---
	withdrawAmount := new(big.Int).Div(depositAmount, big.NewInt(2))
	tx, err = tcContract.StartWithdraw(copyAuth(auths[0]), addrs[3], common.Address{}, withdrawAmount, big.NewInt(1))
	require.NoError(t, err)
	sim.Commit()
	receipt, err = client.TransactionReceipt(context.Background(), tx.Hash())
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt.Status, "startWithdraw failed")

	// Extract withdrawalId from WithdrawStarted event.
	var withdrawalID [32]byte
	for _, log := range receipt.Logs {
		ev, parseErr := iwContract.ParseWithdrawStarted(*log)
		if parseErr == nil {
			withdrawalID = ev.WithdrawalId
			break
		}
	}
	require.NotEqual(t, [32]byte{}, withdrawalID, "WithdrawStarted event not found")

	// --- Compare old (bare estimate) vs new (buffered) gas ---

	// Old behavior: bare estimate via NoSend with GasLimit=0.
	dryRun := copyAuth(auths[1])
	dryRun.NoSend = true
	dryRun.GasLimit = 0
	estTx, err := iwContract.FinalizeWithdraw(dryRun, withdrawalID)
	require.NoError(t, err)
	bareEstimate := estTx.Gas()

	// New behavior: apply 20% buffer.
	bufferedGas := bareEstimate * (100 + gasBufferPercent) / 100

	t.Logf("bare estimate = %d, buffered = %d (+%d%%)", bareEstimate, bufferedGas, gasBufferPercent)
	assert.Greater(t, bufferedGas, bareEstimate, "buffer must increase gas limit")

	// --- Send with buffered gas and verify success ---
	realAuth := copyAuth(auths[1])
	realAuth.GasLimit = bufferedGas
	tx, err = iwContract.FinalizeWithdraw(realAuth, withdrawalID)
	require.NoError(t, err)
	sim.Commit()

	receipt, err = client.TransactionReceipt(context.Background(), tx.Hash())
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt.Status, "finalizeWithdraw with buffered gas should succeed")
	assert.LessOrEqual(t, receipt.GasUsed, bufferedGas,
		"actual gas used should be within the buffered limit")

	t.Logf("gas used = %d, headroom = %d (%.1f%%)",
		receipt.GasUsed, bufferedGas-receipt.GasUsed,
		float64(bufferedGas-receipt.GasUsed)/float64(receipt.GasUsed)*100)
}

// copyAuth creates a shallow copy of TransactOpts so concurrent uses don't race.
func copyAuth(auth *bind.TransactOpts) *bind.TransactOpts {
	cp := *auth
	return &cp
}
