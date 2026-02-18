package checker

import (
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/layer-3/nitewatch/custody"
)

var (
	ErrNoLimitsConfigured      = errors.New("no limits configured for token")
	ErrHourlyLimitExceeded     = errors.New("hourly limit exceeded")
	ErrDailyLimitExceeded      = errors.New("daily limit exceeded")
	ErrUserHourlyLimitExceeded = errors.New("per-user hourly limit exceeded")
	ErrUserDailyLimitExceeded  = errors.New("per-user daily limit exceeded")
	ErrInvalidAmount           = errors.New("amount must be positive")
	ErrInvalidUser             = errors.New("user address must not be zero")
)

type Limit struct {
	Hourly *big.Int
	Daily  *big.Int
}

type Checker struct {
	globalLimits  map[common.Address]Limit
	userOverrides map[common.Address]map[common.Address]Limit
	store         custody.WithdrawalStore
	nowFunc       func() time.Time
}

func New(
	globalLimits map[common.Address]Limit,
	userOverrides map[common.Address]map[common.Address]Limit,
	store custody.WithdrawalStore,
) *Checker {
	return &Checker{
		globalLimits:  globalLimits,
		userOverrides: userOverrides,
		store:         store,
		nowFunc:       time.Now,
	}
}

func (c *Checker) Check(user common.Address, token common.Address, amount *big.Int) error {
	if amount.Sign() <= 0 {
		return ErrInvalidAmount
	}
	if user == (common.Address{}) {
		return ErrInvalidUser
	}

	if err := c.checkGlobalLimits(token, amount); err != nil {
		return err
	}

	if err := c.checkUserLimits(user, token, amount); err != nil {
		return err
	}

	return nil
}

func (c *Checker) checkGlobalLimits(token common.Address, amount *big.Int) error {
	l, ok := c.globalLimits[token]
	if !ok {
		return fmt.Errorf("%w: %s", ErrNoLimitsConfigured, token.Hex())
	}

	now := c.nowFunc()

	if l.Hourly != nil {
		startOfHour := now.Truncate(time.Hour)
		total, err := c.store.GetTotalWithdrawn(token, startOfHour)
		if err != nil {
			return fmt.Errorf("failed to get hourly withdrawn amount: %w", err)
		}
		newTotal := new(big.Int).Add(total, amount)
		if newTotal.Cmp(l.Hourly) > 0 {
			return fmt.Errorf("%w for %s: %s > %s", ErrHourlyLimitExceeded, token.Hex(), newTotal, l.Hourly)
		}
	}

	if l.Daily != nil {
		startOfDay := now.Truncate(24 * time.Hour)
		total, err := c.store.GetTotalWithdrawn(token, startOfDay)
		if err != nil {
			return fmt.Errorf("failed to get daily withdrawn amount: %w", err)
		}
		newTotal := new(big.Int).Add(total, amount)
		if newTotal.Cmp(l.Daily) > 0 {
			return fmt.Errorf("%w for %s: %s > %s", ErrDailyLimitExceeded, token.Hex(), newTotal, l.Daily)
		}
	}

	return nil
}

func (c *Checker) resolveUserLimit(user, token common.Address) *Limit {
	if userTokens, ok := c.userOverrides[user]; ok {
		if l, ok := userTokens[token]; ok {
			return &l
		}
	}
	return nil
}

func (c *Checker) checkUserLimits(user, token common.Address, amount *big.Int) error {
	l := c.resolveUserLimit(user, token)
	if l == nil {
		return nil
	}

	now := c.nowFunc()

	if l.Hourly != nil {
		startOfHour := now.Truncate(time.Hour)
		total, err := c.store.GetTotalWithdrawnByUser(user, token, startOfHour)
		if err != nil {
			return fmt.Errorf("failed to get per-user hourly withdrawn amount: %w", err)
		}
		newTotal := new(big.Int).Add(total, amount)
		if newTotal.Cmp(l.Hourly) > 0 {
			return fmt.Errorf("%w for user %s token %s: %s > %s",
				ErrUserHourlyLimitExceeded, user.Hex(), token.Hex(), newTotal, l.Hourly)
		}
	}

	if l.Daily != nil {
		startOfDay := now.Truncate(24 * time.Hour)
		total, err := c.store.GetTotalWithdrawnByUser(user, token, startOfDay)
		if err != nil {
			return fmt.Errorf("failed to get per-user daily withdrawn amount: %w", err)
		}
		newTotal := new(big.Int).Add(total, amount)
		if newTotal.Cmp(l.Daily) > 0 {
			return fmt.Errorf("%w for user %s token %s: %s > %s",
				ErrUserDailyLimitExceeded, user.Hex(), token.Hex(), newTotal, l.Daily)
		}
	}

	return nil
}

func (c *Checker) Record(w *custody.Withdrawal) error {
	return c.store.Save(w)
}
