package config

import (
	"errors"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Blockchain       BlockchainConfig        `yaml:"blockchain"`
	Limits           LimitsConfig            `yaml:"limits"`
	PerUserOverrides map[string]LimitsConfig `yaml:"per_user_overrides"`
	ListenAddr       string                  `yaml:"listen_addr"`
	DBPath           string                  `yaml:"db_path"`
}

type BlockchainConfig struct {
	RPCURL       string `yaml:"rpc_url"`
	ContractAddr string `yaml:"contract_address"`
	PrivateKey   string `yaml:"private_key"`
}

// LimitsConfig maps token contract addresses to their withdrawal rate limits.
type LimitsConfig map[string]LimitConfig

type LimitConfig struct {
	Hourly string `yaml:"hourly"`
	Daily  string `yaml:"daily"`
}

func (c Config) Validate() error {
	if err := c.Blockchain.Validate(); err != nil {
		return fmt.Errorf("invalid blockchain config: %w", err)
	}
	if len(c.Limits) == 0 {
		return errors.New("at least one token limit must be configured")
	}
	if err := validateLimitsConfig(c.Limits, "limits"); err != nil {
		return err
	}
	for userAddr, tokenLimits := range c.PerUserOverrides {
		if !common.IsHexAddress(userAddr) {
			return fmt.Errorf("invalid user address in per_user_overrides: %s", userAddr)
		}
		if err := validateLimitsConfig(tokenLimits, fmt.Sprintf("per_user_overrides[%s]", userAddr)); err != nil {
			return err
		}
	}
	return nil
}

func validateLimitsConfig(lc LimitsConfig, section string) error {
	for addr, lim := range lc {
		if !common.IsHexAddress(addr) {
			return fmt.Errorf("invalid token address in %s: %s", section, addr)
		}
		if lim.Hourly != "" {
			if _, ok := new(big.Int).SetString(lim.Hourly, 10); !ok {
				return fmt.Errorf("invalid hourly limit for %s in %s: %s", addr, section, lim.Hourly)
			}
		}
		if lim.Daily != "" {
			if _, ok := new(big.Int).SetString(lim.Daily, 10); !ok {
				return fmt.Errorf("invalid daily limit for %s in %s: %s", addr, section, lim.Daily)
			}
		}
	}
	return nil
}

func (c BlockchainConfig) Validate() error {
	if c.RPCURL == "" {
		return errors.New("missing blockchain RPC URL")
	}
	if !strings.HasPrefix(c.RPCURL, "ws://") && !strings.HasPrefix(c.RPCURL, "wss://") {
		return fmt.Errorf("RPC URL must use WebSocket (ws:// or wss://), got: %s", c.RPCURL)
	}
	if !common.IsHexAddress(c.ContractAddr) {
		return fmt.Errorf("invalid contract address: %s", c.ContractAddr)
	}
	if c.PrivateKey == "" {
		return errors.New("missing private key")
	}
	return nil
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config file: %w", err)
	}

	expanded := os.ExpandEnv(string(data))

	var cfg Config
	if err := yaml.Unmarshal([]byte(expanded), &cfg); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	if cfg.DBPath == "" {
		cfg.DBPath = "nitewatch.db"
	}

	if cfg.ListenAddr == "" {
		cfg.ListenAddr = ":8080"
	}

	return &cfg, nil
}
