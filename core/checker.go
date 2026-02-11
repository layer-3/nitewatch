package core

import (
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	nw "github.com/layer-3/nitewatch"
)

var (
	ErrNoLimitsConfigured  = errors.New("no limits configured for token")
	ErrHourlyLimitExceeded = errors.New("hourly limit exceeded")
	ErrDailyLimitExceeded  = errors.New("daily limit exceeded")
)

// LimitConfig defines the withdrawal constraints for a token (YAML friendly).
type LimitConfig struct {
	Hourly string `yaml:"hourly"`
	Daily  string `yaml:"daily"`
}

// Config maps token addresses to their limits.
type Config struct {
	Limits map[string]LimitConfig `yaml:"limits"`
}

// limit defines the internal parsed withdrawal constraints.
type limit struct {
	Hourly *big.Int
	Daily  *big.Int
}

// Checker manages withdrawal limits using a database store.
type Checker struct {
	limits  map[common.Address]limit
	store   nw.WithdrawalStore
	nowFunc func() time.Time
}

// NewChecker creates a new limit checker with the provided configuration and store.
func NewChecker(cfg Config, store nw.WithdrawalStore) (*Checker, error) {
	limits := make(map[common.Address]limit)
	for addrStr, conf := range cfg.Limits {
		if !common.IsHexAddress(addrStr) {
			return nil, fmt.Errorf("invalid address in config: %s", addrStr)
		}
		addr := common.HexToAddress(addrStr)

		l := limit{}
		if conf.Hourly != "" {
			val, ok := new(big.Int).SetString(conf.Hourly, 10)
			if !ok {
				return nil, fmt.Errorf("invalid hourly limit for %s: %s", addrStr, conf.Hourly)
			}
			l.Hourly = val
		}
		if conf.Daily != "" {
			val, ok := new(big.Int).SetString(conf.Daily, 10)
			if !ok {
				return nil, fmt.Errorf("invalid daily limit for %s: %s", addrStr, conf.Daily)
			}
			l.Daily = val
		}
		limits[addr] = l
	}

	return &Checker{
		limits:  limits,
		store:   store,
		nowFunc: time.Now,
	}, nil
}

// Check verifies if a withdrawal amount is within limits for the given token.
// It queries the store for total withdrawn amounts in the current hour and day.
func (c *Checker) Check(token common.Address, amount *big.Int) error {
	l, ok := c.limits[token]
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

// Record persists the withdrawal event to the store.
func (c *Checker) Record(w *nw.Withdrawal) error {
	return c.store.Save(w)
}
