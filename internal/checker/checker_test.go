package checker

import (
	"errors"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/layer-3/nitewatch/custody"
)

type mockStore struct {
	withdrawals []*custody.Withdrawal
	err         error
}

func (m *mockStore) Save(w *custody.Withdrawal) error {
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

func (m *mockStore) GetTotalWithdrawnByUser(user common.Address, token common.Address, since time.Time) (*big.Int, error) {
	if m.err != nil {
		return nil, m.err
	}
	total := new(big.Int)
	for _, w := range m.withdrawals {
		if w.User == user && w.Token == token && !w.Timestamp.Before(since) {
			total.Add(total, w.Amount)
		}
	}
	return total, nil
}

var (
	tokenA = common.HexToAddress("0xAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	tokenB = common.HexToAddress("0xBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB")
	userA  = common.HexToAddress("0x1111111111111111111111111111111111111111")
	userB  = common.HexToAddress("0x2222222222222222222222222222222222222222")
)

func globalLimits(token common.Address, hourly, daily *big.Int) map[common.Address]Limit {
	return map[common.Address]Limit{
		token: {Hourly: hourly, Daily: daily},
	}
}

func TestNew_Valid(t *testing.T) {
	c := New(globalLimits(tokenA, big.NewInt(1000), big.NewInt(5000)), nil, &mockStore{})
	require.NotNil(t, c)
}

func TestCheck_SanityChecks(t *testing.T) {
	c := New(globalLimits(tokenA, big.NewInt(1000), big.NewInt(5000)), nil, &mockStore{})

	t.Run("zero amount", func(t *testing.T) {
		err := c.Check(userA, tokenA, big.NewInt(0))
		require.ErrorIs(t, err, ErrInvalidAmount)
	})

	t.Run("negative amount", func(t *testing.T) {
		err := c.Check(userA, tokenA, big.NewInt(-1))
		require.ErrorIs(t, err, ErrInvalidAmount)
	})

	t.Run("zero user address", func(t *testing.T) {
		err := c.Check(common.Address{}, tokenA, big.NewInt(100))
		require.ErrorIs(t, err, ErrInvalidUser)
	})
}

func TestCheck_NoLimitsConfigured(t *testing.T) {
	c := New(globalLimits(tokenA, big.NewInt(1000), nil), nil, &mockStore{})

	err := c.Check(userA, tokenB, big.NewInt(100))
	require.ErrorIs(t, err, ErrNoLimitsConfigured)
}

func TestCheck_UnderHourlyLimit(t *testing.T) {
	c := New(globalLimits(tokenA, big.NewInt(1000), big.NewInt(5000)), nil, &mockStore{})
	c.nowFunc = func() time.Time {
		return time.Date(2025, 1, 1, 12, 30, 0, 0, time.UTC)
	}

	require.NoError(t, c.Check(userA, tokenA, big.NewInt(500)))
}

func TestCheck_ExactHourlyLimit(t *testing.T) {
	c := New(globalLimits(tokenA, big.NewInt(1000), big.NewInt(5000)), nil, &mockStore{})
	c.nowFunc = func() time.Time {
		return time.Date(2025, 1, 1, 12, 30, 0, 0, time.UTC)
	}

	require.NoError(t, c.Check(userA, tokenA, big.NewInt(1000)))
}

func TestCheck_ExceedHourlyLimit(t *testing.T) {
	now := time.Date(2025, 1, 1, 12, 30, 0, 0, time.UTC)
	store := &mockStore{
		withdrawals: []*custody.Withdrawal{
			{Token: tokenA, User: userA, Amount: big.NewInt(800), Timestamp: now.Add(-10 * time.Minute)},
		},
	}
	c := New(globalLimits(tokenA, big.NewInt(1000), big.NewInt(5000)), nil, store)
	c.nowFunc = func() time.Time { return now }

	err := c.Check(userA, tokenA, big.NewInt(300))
	require.ErrorIs(t, err, ErrHourlyLimitExceeded)
}

func TestCheck_ExceedDailyLimit(t *testing.T) {
	now := time.Date(2025, 1, 1, 12, 30, 0, 0, time.UTC)
	store := &mockStore{
		withdrawals: []*custody.Withdrawal{
			{Token: tokenA, User: userA, Amount: big.NewInt(4500), Timestamp: now.Add(-3 * time.Hour)},
		},
	}
	c := New(globalLimits(tokenA, big.NewInt(1000), big.NewInt(5000)), nil, store)
	c.nowFunc = func() time.Time { return now }

	err := c.Check(userA, tokenA, big.NewInt(600))
	require.ErrorIs(t, err, ErrDailyLimitExceeded)
}

func TestCheck_PreviousHourNotCounted(t *testing.T) {
	now := time.Date(2025, 1, 1, 13, 5, 0, 0, time.UTC)
	store := &mockStore{
		withdrawals: []*custody.Withdrawal{
			{Token: tokenA, User: userA, Amount: big.NewInt(900), Timestamp: time.Date(2025, 1, 1, 12, 50, 0, 0, time.UTC)},
		},
	}
	c := New(globalLimits(tokenA, big.NewInt(1000), big.NewInt(5000)), nil, store)
	c.nowFunc = func() time.Time { return now }

	require.NoError(t, c.Check(userA, tokenA, big.NewInt(900)))
}

func TestCheck_HourlyOnlyConfig(t *testing.T) {
	c := New(globalLimits(tokenA, big.NewInt(1000), nil), nil, &mockStore{})
	c.nowFunc = func() time.Time {
		return time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC)
	}

	require.NoError(t, c.Check(userA, tokenA, big.NewInt(999)))
}

func TestCheck_DailyOnlyConfig(t *testing.T) {
	c := New(globalLimits(tokenA, nil, big.NewInt(5000)), nil, &mockStore{})
	c.nowFunc = func() time.Time {
		return time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC)
	}

	require.NoError(t, c.Check(userA, tokenA, big.NewInt(4000)))
}

func TestCheck_StoreError(t *testing.T) {
	store := &mockStore{err: errors.New("db connection lost")}
	c := New(globalLimits(tokenA, big.NewInt(1000), big.NewInt(5000)), nil, store)
	c.nowFunc = func() time.Time {
		return time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC)
	}

	err := c.Check(userA, tokenA, big.NewInt(100))
	require.Error(t, err)
}

func TestRecord(t *testing.T) {
	store := &mockStore{}
	c := New(globalLimits(tokenA, big.NewInt(1000), nil), nil, store)

	w := &custody.Withdrawal{
		WithdrawalID: [32]byte{1},
		User:         userA,
		Token:        tokenA,
		Amount:       big.NewInt(500),
		Timestamp:    time.Now(),
	}
	require.NoError(t, c.Record(w))
	require.Len(t, store.withdrawals, 1)
}

// --- Per-user limit tests ---

func TestCheck_PerUserOverrideTakesPrecedence(t *testing.T) {
	now := time.Date(2025, 1, 1, 12, 30, 0, 0, time.UTC)
	store := &mockStore{
		withdrawals: []*custody.Withdrawal{
			{Token: tokenA, User: userA, Amount: big.NewInt(800), Timestamp: now.Add(-10 * time.Minute)},
		},
	}

	global := globalLimits(tokenA, big.NewInt(10000), big.NewInt(50000))
	overrides := map[common.Address]map[common.Address]Limit{
		userA: {tokenA: {Hourly: big.NewInt(1000), Daily: big.NewInt(5000)}},
	}

	c := New(global, overrides, store)
	c.nowFunc = func() time.Time { return now }

	// userA has override of 1000 hourly; 800 + 150 = 950 < 1000 → pass
	require.NoError(t, c.Check(userA, tokenA, big.NewInt(150)))

	// userA: 800 + 250 = 1050 > 1000 → per-user hourly exceeded
	err := c.Check(userA, tokenA, big.NewInt(250))
	require.ErrorIs(t, err, ErrUserHourlyLimitExceeded)
}

func TestCheck_PerUserLimitIndependentOfGlobal(t *testing.T) {
	now := time.Date(2025, 1, 1, 12, 30, 0, 0, time.UTC)

	store := &mockStore{
		withdrawals: []*custody.Withdrawal{
			{Token: tokenA, User: userA, Amount: big.NewInt(400), Timestamp: now.Add(-10 * time.Minute)},
			{Token: tokenA, User: userB, Amount: big.NewInt(400), Timestamp: now.Add(-5 * time.Minute)},
		},
	}

	global := globalLimits(tokenA, big.NewInt(1000), nil)
	overrides := map[common.Address]map[common.Address]Limit{
		userA: {tokenA: {Hourly: big.NewInt(500)}},
		userB: {tokenA: {Hourly: big.NewInt(500)}},
	}

	c := New(global, overrides, store)
	c.nowFunc = func() time.Time { return now }

	// userA: 400 + 200 = 600 > per-user 500 → blocked by per-user limit
	err := c.Check(userA, tokenA, big.NewInt(200))
	require.ErrorIs(t, err, ErrUserHourlyLimitExceeded)

	// global: 800 + 250 = 1050 > 1000 → blocked by global limit
	err = c.Check(userB, tokenA, big.NewInt(250))
	require.Error(t, err)
}

func TestCheck_PerUserDailyLimitExceeded(t *testing.T) {
	now := time.Date(2025, 1, 1, 18, 30, 0, 0, time.UTC)
	store := &mockStore{
		withdrawals: []*custody.Withdrawal{
			{Token: tokenA, User: userA, Amount: big.NewInt(1500), Timestamp: now.Add(-6 * time.Hour)},
			{Token: tokenA, User: userA, Amount: big.NewInt(400), Timestamp: now.Add(-2 * time.Hour)},
		},
	}

	global := globalLimits(tokenA, big.NewInt(10000), big.NewInt(50000))
	overrides := map[common.Address]map[common.Address]Limit{
		userA: {tokenA: {Hourly: big.NewInt(5000), Daily: big.NewInt(2000)}},
	}

	c := New(global, overrides, store)
	c.nowFunc = func() time.Time { return now }

	// userA already withdrew 1500+400=1900 today; 1900+200=2100 > daily limit 2000
	err := c.Check(userA, tokenA, big.NewInt(200))
	require.ErrorIs(t, err, ErrUserDailyLimitExceeded)
}

func TestCheck_NoPerUserLimitsConfigured(t *testing.T) {
	c := New(globalLimits(tokenA, big.NewInt(1000), big.NewInt(5000)), nil, &mockStore{})
	c.nowFunc = func() time.Time {
		return time.Date(2025, 1, 1, 12, 0, 0, 0, time.UTC)
	}

	require.NoError(t, c.Check(userA, tokenA, big.NewInt(500)))
}
