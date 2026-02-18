package store

import (
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/layer-3/nitewatch/custody"
)

func newTestAdapter(t *testing.T) *Adapter {
	t.Helper()
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		Logger: logger.Discard,
	})
	require.NoError(t, err)
	adapter, err := NewAdapter(db)
	require.NoError(t, err)
	return adapter
}

var (
	tokenA = common.HexToAddress("0xAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	tokenB = common.HexToAddress("0xBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB")
	user   = common.HexToAddress("0x1111111111111111111111111111111111111111")
)

func TestSave(t *testing.T) {
	a := newTestAdapter(t)

	w := &custody.Withdrawal{
		WithdrawalID: [32]byte{1},
		User:         user,
		Token:        tokenA,
		Amount:       big.NewInt(1000),
		BlockNumber:  42,
		TxHash:       common.HexToHash("0xdeadbeef"),
		Timestamp:    time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC),
	}

	require.NoError(t, a.Save(w))

	total, err := a.GetTotalWithdrawn(tokenA, time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC))
	require.NoError(t, err)
	require.Equal(t, "1000", total.String())
}

func TestSave_DuplicateWithdrawalID(t *testing.T) {
	a := newTestAdapter(t)

	w := &custody.Withdrawal{
		WithdrawalID: [32]byte{1},
		User:         user,
		Token:        tokenA,
		Amount:       big.NewInt(1000),
		Timestamp:    time.Now(),
	}

	require.NoError(t, a.Save(w))
	require.Error(t, a.Save(w))
}

func TestGetTotalWithdrawn_Empty(t *testing.T) {
	a := newTestAdapter(t)

	total, err := a.GetTotalWithdrawn(tokenA, time.Time{})
	require.NoError(t, err)
	require.Equal(t, 0, total.Sign())
}

func TestGetTotalWithdrawn_TimeFilter(t *testing.T) {
	a := newTestAdapter(t)

	base := time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC)

	withdrawals := []*custody.Withdrawal{
		{WithdrawalID: [32]byte{1}, User: user, Token: tokenA, Amount: big.NewInt(100), Timestamp: base.Add(-2 * time.Hour)},
		{WithdrawalID: [32]byte{2}, User: user, Token: tokenA, Amount: big.NewInt(200), Timestamp: base.Add(-30 * time.Minute)},
		{WithdrawalID: [32]byte{3}, User: user, Token: tokenA, Amount: big.NewInt(300), Timestamp: base.Add(10 * time.Minute)},
	}
	for _, w := range withdrawals {
		require.NoError(t, a.Save(w))
	}

	total, err := a.GetTotalWithdrawn(tokenA, base)
	require.NoError(t, err)
	require.Equal(t, "300", total.String())

	total, err = a.GetTotalWithdrawn(tokenA, base.Add(-1*time.Hour))
	require.NoError(t, err)
	require.Equal(t, "500", total.String())

	total, err = a.GetTotalWithdrawn(tokenA, base.Add(-3*time.Hour))
	require.NoError(t, err)
	require.Equal(t, "600", total.String())
}

func TestGetTotalWithdrawn_TokenFilter(t *testing.T) {
	a := newTestAdapter(t)

	base := time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC)

	withdrawals := []*custody.Withdrawal{
		{WithdrawalID: [32]byte{1}, User: user, Token: tokenA, Amount: big.NewInt(100), Timestamp: base},
		{WithdrawalID: [32]byte{2}, User: user, Token: tokenB, Amount: big.NewInt(200), Timestamp: base},
		{WithdrawalID: [32]byte{3}, User: user, Token: tokenA, Amount: big.NewInt(300), Timestamp: base},
	}
	for _, w := range withdrawals {
		require.NoError(t, a.Save(w))
	}

	totalA, err := a.GetTotalWithdrawn(tokenA, base.Add(-time.Hour))
	require.NoError(t, err)
	require.Equal(t, "400", totalA.String())

	totalB, err := a.GetTotalWithdrawn(tokenB, base.Add(-time.Hour))
	require.NoError(t, err)
	require.Equal(t, "200", totalB.String())
}

func TestGetTotalWithdrawnByUser(t *testing.T) {
	a := newTestAdapter(t)
	base := time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC)

	userB := common.HexToAddress("0x2222222222222222222222222222222222222222")

	withdrawals := []*custody.Withdrawal{
		{WithdrawalID: [32]byte{1}, User: user, Token: tokenA, Amount: big.NewInt(100), Timestamp: base},
		{WithdrawalID: [32]byte{2}, User: user, Token: tokenA, Amount: big.NewInt(200), Timestamp: base},
		{WithdrawalID: [32]byte{3}, User: userB, Token: tokenA, Amount: big.NewInt(300), Timestamp: base},
		{WithdrawalID: [32]byte{4}, User: user, Token: tokenB, Amount: big.NewInt(400), Timestamp: base},
	}
	for _, w := range withdrawals {
		require.NoError(t, a.Save(w))
	}

	// user + tokenA = 100 + 200 = 300
	total, err := a.GetTotalWithdrawnByUser(user, tokenA, base.Add(-time.Hour))
	require.NoError(t, err)
	require.Equal(t, "300", total.String())

	// userB + tokenA = 300
	total, err = a.GetTotalWithdrawnByUser(userB, tokenA, base.Add(-time.Hour))
	require.NoError(t, err)
	require.Equal(t, "300", total.String())

	// user + tokenB = 400
	total, err = a.GetTotalWithdrawnByUser(user, tokenB, base.Add(-time.Hour))
	require.NoError(t, err)
	require.Equal(t, "400", total.String())

	// userB + tokenB = 0 (no withdrawals)
	total, err = a.GetTotalWithdrawnByUser(userB, tokenB, base.Add(-time.Hour))
	require.NoError(t, err)
	require.Equal(t, 0, total.Sign())
}

func TestGetTotalWithdrawn_LargeAmounts(t *testing.T) {
	a := newTestAdapter(t)
	base := time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC)

	bigAmount, _ := new(big.Int).SetString("999999999999999999999999999999", 10)

	w := &custody.Withdrawal{
		WithdrawalID: [32]byte{1},
		User:         user,
		Token:        tokenA,
		Amount:       bigAmount,
		Timestamp:    base,
	}
	require.NoError(t, a.Save(w))

	total, err := a.GetTotalWithdrawn(tokenA, base.Add(-time.Hour))
	require.NoError(t, err)
	require.Equal(t, bigAmount.String(), total.String())
}
