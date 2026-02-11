package store

import (
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	nw "github.com/layer-3/nitewatch"
)

func newTestAdapter(t *testing.T) *Adapter {
	t.Helper()
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		Logger: logger.Discard,
	})
	if err != nil {
		t.Fatalf("failed to open in-memory db: %v", err)
	}
	adapter, err := NewAdapter(db)
	if err != nil {
		t.Fatalf("failed to create adapter: %v", err)
	}
	return adapter
}

var (
	tokenA = common.HexToAddress("0xAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	tokenB = common.HexToAddress("0xBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB")
	user   = common.HexToAddress("0x1111111111111111111111111111111111111111")
)

func TestSave(t *testing.T) {
	a := newTestAdapter(t)

	w := &nw.Withdrawal{
		WithdrawalID: [32]byte{1},
		User:         user,
		Token:        tokenA,
		Amount:       big.NewInt(1000),
		BlockNumber:  42,
		TxHash:       common.HexToHash("0xdeadbeef"),
		Timestamp:    time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC),
	}

	if err := a.Save(w); err != nil {
		t.Fatalf("Save failed: %v", err)
	}

	// Verify it was persisted
	total, err := a.GetTotalWithdrawn(tokenA, time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC))
	if err != nil {
		t.Fatalf("GetTotalWithdrawn failed: %v", err)
	}
	if total.Cmp(big.NewInt(1000)) != 0 {
		t.Fatalf("expected total 1000, got %s", total)
	}
}

func TestSave_DuplicateWithdrawalID(t *testing.T) {
	a := newTestAdapter(t)

	w := &nw.Withdrawal{
		WithdrawalID: [32]byte{1},
		User:         user,
		Token:        tokenA,
		Amount:       big.NewInt(1000),
		Timestamp:    time.Now(),
	}

	if err := a.Save(w); err != nil {
		t.Fatalf("first Save failed: %v", err)
	}
	if err := a.Save(w); err == nil {
		t.Fatal("expected error on duplicate withdrawal ID")
	}
}

func TestGetTotalWithdrawn_Empty(t *testing.T) {
	a := newTestAdapter(t)

	total, err := a.GetTotalWithdrawn(tokenA, time.Time{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total.Sign() != 0 {
		t.Fatalf("expected zero, got %s", total)
	}
}

func TestGetTotalWithdrawn_TimeFilter(t *testing.T) {
	a := newTestAdapter(t)

	base := time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC)

	withdrawals := []*nw.Withdrawal{
		{WithdrawalID: [32]byte{1}, User: user, Token: tokenA, Amount: big.NewInt(100), Timestamp: base.Add(-2 * time.Hour)},
		{WithdrawalID: [32]byte{2}, User: user, Token: tokenA, Amount: big.NewInt(200), Timestamp: base.Add(-30 * time.Minute)},
		{WithdrawalID: [32]byte{3}, User: user, Token: tokenA, Amount: big.NewInt(300), Timestamp: base.Add(10 * time.Minute)},
	}
	for _, w := range withdrawals {
		if err := a.Save(w); err != nil {
			t.Fatalf("Save failed: %v", err)
		}
	}

	// Since base: should include w2 (at base-30m? No, base-30m < base) and w3
	// Actually base-30m is before base, so only w3 is >= base
	// Let me check: since = base = 12:00
	// w1 = 10:00 (before), w2 = 11:30 (before), w3 = 12:10 (after)
	total, err := a.GetTotalWithdrawn(tokenA, base)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total.Cmp(big.NewInt(300)) != 0 {
		t.Fatalf("expected 300, got %s", total)
	}

	// Since 11:00: should include w2 and w3
	total, err = a.GetTotalWithdrawn(tokenA, base.Add(-1*time.Hour))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total.Cmp(big.NewInt(500)) != 0 {
		t.Fatalf("expected 500, got %s", total)
	}

	// Since beginning: all three
	total, err = a.GetTotalWithdrawn(tokenA, base.Add(-3*time.Hour))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total.Cmp(big.NewInt(600)) != 0 {
		t.Fatalf("expected 600, got %s", total)
	}
}

func TestGetTotalWithdrawn_TokenFilter(t *testing.T) {
	a := newTestAdapter(t)

	base := time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC)

	withdrawals := []*nw.Withdrawal{
		{WithdrawalID: [32]byte{1}, User: user, Token: tokenA, Amount: big.NewInt(100), Timestamp: base},
		{WithdrawalID: [32]byte{2}, User: user, Token: tokenB, Amount: big.NewInt(200), Timestamp: base},
		{WithdrawalID: [32]byte{3}, User: user, Token: tokenA, Amount: big.NewInt(300), Timestamp: base},
	}
	for _, w := range withdrawals {
		if err := a.Save(w); err != nil {
			t.Fatalf("Save failed: %v", err)
		}
	}

	totalA, err := a.GetTotalWithdrawn(tokenA, base.Add(-time.Hour))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if totalA.Cmp(big.NewInt(400)) != 0 {
		t.Fatalf("expected 400 for tokenA, got %s", totalA)
	}

	totalB, err := a.GetTotalWithdrawn(tokenB, base.Add(-time.Hour))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if totalB.Cmp(big.NewInt(200)) != 0 {
		t.Fatalf("expected 200 for tokenB, got %s", totalB)
	}
}

func TestGetTotalWithdrawn_LargeAmounts(t *testing.T) {
	a := newTestAdapter(t)
	base := time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC)

	// Use amounts larger than uint64
	bigAmount, _ := new(big.Int).SetString("999999999999999999999999999999", 10)

	w := &nw.Withdrawal{
		WithdrawalID: [32]byte{1},
		User:         user,
		Token:        tokenA,
		Amount:       bigAmount,
		Timestamp:    base,
	}
	if err := a.Save(w); err != nil {
		t.Fatalf("Save failed: %v", err)
	}

	total, err := a.GetTotalWithdrawn(tokenA, base.Add(-time.Hour))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if total.Cmp(bigAmount) != 0 {
		t.Fatalf("expected %s, got %s", bigAmount, total)
	}
}
