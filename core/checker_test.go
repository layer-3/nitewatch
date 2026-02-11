package core

import (
	"errors"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	nw "github.com/layer-3/nitewatch"
)

// mockStore is an in-memory WithdrawalStore for testing.
type mockStore struct {
	withdrawals []*nw.Withdrawal
	err         error // if set, GetTotalWithdrawn returns this error
}

func (m *mockStore) Save(w *nw.Withdrawal) error {
	m.withdrawals = append(m.withdrawals, w)
	return nil
}

func (m *mockStore) GetTotalWithdrawn(token common.Address, since time.Time) (*big.Int, error) {
	if m.err != nil {
		return nil, m.err
	}
	total := new(big.Int)
	for _, w := range m.withdrawals {
		if w.Token == token && !w.Timestamp.Before(since) {
			total.Add(total, w.Amount)
		}
	}
	return total, nil
}

var (
	tokenA = common.HexToAddress("0xAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	tokenB = common.HexToAddress("0xBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB")
	user   = common.HexToAddress("0x1111111111111111111111111111111111111111")
)

func TestNewChecker_ValidConfig(t *testing.T) {
	cfg := Config{
		Limits: map[string]LimitConfig{
			tokenA.Hex(): {Hourly: "1000", Daily: "5000"},
		},
	}
	c, err := NewChecker(cfg, &mockStore{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if c == nil {
		t.Fatal("expected non-nil checker")
	}
}

func TestNewChecker_InvalidAddress(t *testing.T) {
	cfg := Config{
		Limits: map[string]LimitConfig{
			"not-an-address": {Hourly: "1000"},
		},
	}
	_, err := NewChecker(cfg, &mockStore{})
	if err == nil {
		t.Fatal("expected error for invalid address")
	}
}

func TestNewChecker_InvalidHourlyLimit(t *testing.T) {
	cfg := Config{
		Limits: map[string]LimitConfig{
			tokenA.Hex(): {Hourly: "not-a-number"},
		},
	}
	_, err := NewChecker(cfg, &mockStore{})
	if err == nil {
		t.Fatal("expected error for invalid hourly limit")
	}
}

func TestNewChecker_InvalidDailyLimit(t *testing.T) {
	cfg := Config{
		Limits: map[string]LimitConfig{
			tokenA.Hex(): {Daily: "xyz"},
		},
	}
	_, err := NewChecker(cfg, &mockStore{})
	if err == nil {
		t.Fatal("expected error for invalid daily limit")
	}
}

func TestCheck_NoLimitsConfigured(t *testing.T) {
	cfg := Config{
		Limits: map[string]LimitConfig{
			tokenA.Hex(): {Hourly: "1000"},
		},
	}
	c, err := NewChecker(cfg, &mockStore{})
	if err != nil {
		t.Fatal(err)
	}

	err = c.Check(tokenB, big.NewInt(100))
	if !errors.Is(err, ErrNoLimitsConfigured) {
		t.Fatalf("expected ErrNoLimitsConfigured, got: %v", err)
	}
}

func TestCheck_UnderHourlyLimit(t *testing.T) {
	store := &mockStore{}
	cfg := Config{
		Limits: map[string]LimitConfig{
			tokenA.Hex(): {Hourly: "1000", Daily: "5000"},
		},
	}
	c, err := NewChecker(cfg, store)
	if err != nil {
		t.Fatal(err)
	}
	c.nowFunc = func() time.Time {
		return time.Date(2025, 1, 1, 12, 30, 0, 0, time.UTC)
	}

	if err := c.Check(tokenA, big.NewInt(500)); err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
}

func TestCheck_ExactHourlyLimit(t *testing.T) {
	store := &mockStore{}
	cfg := Config{
		Limits: map[string]LimitConfig{
			tokenA.Hex(): {Hourly: "1000", Daily: "5000"},
		},
	}
	c, err := NewChecker(cfg, store)
	if err != nil {
		t.Fatal(err)
	}
	c.nowFunc = func() time.Time {
		return time.Date(2025, 1, 1, 12, 30, 0, 0, time.UTC)
	}

	// Exact limit should pass (not exceed)
	if err := c.Check(tokenA, big.NewInt(1000)); err != nil {
		t.Fatalf("expected no error at exact limit, got: %v", err)
	}
}

func TestCheck_ExceedHourlyLimit(t *testing.T) {
	now := time.Date(2025, 1, 1, 12, 30, 0, 0, time.UTC)
	store := &mockStore{
		withdrawals: []*nw.Withdrawal{
			{Token: tokenA, Amount: big.NewInt(800), Timestamp: now.Add(-10 * time.Minute)},
		},
	}
	cfg := Config{
		Limits: map[string]LimitConfig{
			tokenA.Hex(): {Hourly: "1000", Daily: "5000"},
		},
	}
	c, err := NewChecker(cfg, store)
	if err != nil {
		t.Fatal(err)
	}
	c.nowFunc = func() time.Time { return now }

	err = c.Check(tokenA, big.NewInt(300))
	if !errors.Is(err, ErrHourlyLimitExceeded) {
		t.Fatalf("expected ErrHourlyLimitExceeded, got: %v", err)
	}
}

func TestCheck_ExceedDailyLimit(t *testing.T) {
	now := time.Date(2025, 1, 1, 12, 30, 0, 0, time.UTC)
	store := &mockStore{
		withdrawals: []*nw.Withdrawal{
			// Old withdrawal within the day but outside the hour
			{Token: tokenA, Amount: big.NewInt(4500), Timestamp: now.Add(-3 * time.Hour)},
		},
	}
	cfg := Config{
		Limits: map[string]LimitConfig{
			tokenA.Hex(): {Hourly: "1000", Daily: "5000"},
		},
	}
	c, err := NewChecker(cfg, store)
	if err != nil {
		t.Fatal(err)
	}
	c.nowFunc = func() time.Time { return now }

	err = c.Check(tokenA, big.NewInt(600))
	if !errors.Is(err, ErrDailyLimitExceeded) {
		t.Fatalf("expected ErrDailyLimitExceeded, got: %v", err)
	}
}

func TestCheck_PreviousHourNotCounted(t *testing.T) {
	now := time.Date(2025, 1, 1, 13, 5, 0, 0, time.UTC)
	store := &mockStore{
		withdrawals: []*nw.Withdrawal{
			// Withdrawal from the previous hour
			{Token: tokenA, Amount: big.NewInt(900), Timestamp: time.Date(2025, 1, 1, 12, 50, 0, 0, time.UTC)},
		},
	}
	cfg := Config{
		Limits: map[string]LimitConfig{
			tokenA.Hex(): {Hourly: "1000", Daily: "5000"},
		},
	}
	c, err := NewChecker(cfg, store)
	if err != nil {
		t.Fatal(err)
	}
	c.nowFunc = func() time.Time { return now }

	// Previous hour withdrawal shouldn't count towards hourly limit
	if err := c.Check(tokenA, big.NewInt(900)); err != nil {
		t.Fatalf("expected no error (previous hour), got: %v", err)
	}
}

func TestCheck_HourlyOnlyConfig(t *testing.T) {
	store := &mockStore{}
	cfg := Config{
		Limits: map[string]LimitConfig{
			tokenA.Hex(): {Hourly: "1000"},
		},
	}
	c, err := NewChecker(cfg, store)
	if err != nil {
		t.Fatal(err)
	}
	c.nowFunc = func() time.Time {
		return time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC)
	}

	// Without daily limit, large amount within hourly should pass
	if err := c.Check(tokenA, big.NewInt(999)); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestCheck_DailyOnlyConfig(t *testing.T) {
	store := &mockStore{}
	cfg := Config{
		Limits: map[string]LimitConfig{
			tokenA.Hex(): {Daily: "5000"},
		},
	}
	c, err := NewChecker(cfg, store)
	if err != nil {
		t.Fatal(err)
	}
	c.nowFunc = func() time.Time {
		return time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC)
	}

	// Without hourly limit, this should pass
	if err := c.Check(tokenA, big.NewInt(4000)); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestCheck_StoreError(t *testing.T) {
	store := &mockStore{err: errors.New("db connection lost")}
	cfg := Config{
		Limits: map[string]LimitConfig{
			tokenA.Hex(): {Hourly: "1000", Daily: "5000"},
		},
	}
	c, err := NewChecker(cfg, store)
	if err != nil {
		t.Fatal(err)
	}
	c.nowFunc = func() time.Time {
		return time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC)
	}

	err = c.Check(tokenA, big.NewInt(100))
	if err == nil {
		t.Fatal("expected error from store")
	}
}

func TestRecord(t *testing.T) {
	store := &mockStore{}
	cfg := Config{
		Limits: map[string]LimitConfig{
			tokenA.Hex(): {Hourly: "1000"},
		},
	}
	c, err := NewChecker(cfg, store)
	if err != nil {
		t.Fatal(err)
	}

	w := &nw.Withdrawal{
		WithdrawalID: [32]byte{1},
		User:         user,
		Token:        tokenA,
		Amount:       big.NewInt(500),
		Timestamp:    time.Now(),
	}
	if err := c.Record(w); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(store.withdrawals) != 1 {
		t.Fatalf("expected 1 withdrawal in store, got %d", len(store.withdrawals))
	}
}
