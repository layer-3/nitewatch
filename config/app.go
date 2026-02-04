package config

import (
	"fmt"
	"time"

	"github.com/layer-3/nitewatch/core/log"
)

// AppConfig holds all application configuration
type App struct {
	Port     uint         `env:"APP_PORT" env-default:"8080"`
	Mode     string       `env:"APP_MODE" env-default:"development"`
	LogLevel log.LogLevel `env:"APP_LOG_LEVEL" env-default:"debug"`
	Auth     Auth
	Database Database
	Redis    Redis
	Agent    Agent
	Security Security
}

// Security holds security policy configuration
type Security struct {
	DefaultUserHourlyLimit string `env:"SEC_USER_HOURLY_LIMIT" env-default:"1000000000000000000"`   // 1 ETH (wei)
	DefaultUserDailyLimit  string `env:"SEC_USER_DAILY_LIMIT"  env-default:"10000000000000000000"`  // 10 ETH (wei)
	GlobalHourlyLimit      string `env:"SEC_GLOBAL_HOURLY_LIMIT" env-default:"100000000000000000000"` // 100 ETH (wei)
	GlobalDailyLimit       string `env:"SEC_GLOBAL_DAILY_LIMIT"  env-default:"1000000000000000000000"` // 1000 ETH (wei)
}

// Auth holds authentication configuration
type Auth struct {
	JWTSecret       string        `env:"AUTH_JWT_SECRET" env-required:"true"`
	AccessDuration  time.Duration `env:"AUTH_ACCESS_DURATION" env-default:"5m"`
	RefreshDuration time.Duration `env:"AUTH_REFRESH_DURATION" env-default:"168h"` // 7 days
	ResendAPIKey    string        `env:"AUTH_RESEND_API_KEY"`
	FromEmail       string        `env:"AUTH_FROM_EMAIL" env-default:"noreply@nitewatch.app"`
}

// Agent holds AI service configuration
type Agent struct {
	Provider string `env:"AGENT_PROVIDER" env-default:"gemini"`
	APIKey   string `env:"AGENT_API_KEY"`
	Model    string `env:"AGENT_MODEL" env-default:"gemini-2.5-flash"`
}

// LoadAppConfig loads the application configuration from environment variables
func LoadAppConfig() (*App, error) {
	cfg := &App{}
	if err := Load(cfg); err != nil {
		return nil, fmt.Errorf("failed to load app config: %w", err)
	}

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	if err := cfg.ConfigureLogger(); err != nil {
		return nil, fmt.Errorf("failed to configure logger: %w", err)
	}

	return cfg, nil
}

// Validate checks the configuration for validity
func (c *App) Validate() error {
	switch c.Mode {
	case "development", "production", "test":
		return nil
	default:
		return fmt.Errorf("invalid APP_MODE '%s', must be one of: development, production, test", c.Mode)
	}
}

// ConfigureLogger configures the global logger based on the App configuration.
func (c *App) ConfigureLogger() error {
	logLevel := c.GetLevel()

	if err := log.SetLevel(logLevel); err != nil {
		return fmt.Errorf("invalid log level '%s': %w", logLevel, err)
	}

	// Enable console logger for development (default mode)
	switch c.Mode {
	case "production", "test":
		// no console logger
	default:
		log.EnableConsoleLogger()
	}

	return nil
}

// GetGinMode returns the appropriate Gin mode string based on the App configuration.
func (c *App) GetGinMode() string {
	switch c.Mode {
	case "production", "test":
		return "release"
	default:
		return "debug"
	}
}

// GetLevel returns the configured log level or a default based on the mode
func (c *App) GetLevel() log.LogLevel {
	if c.LogLevel != "" {
		return c.LogLevel
	}

	switch c.Mode {
	case "production":
		return log.LevelWarn
	case "test":
		return log.LevelInfo
	default:
		return log.LevelDebug
	}
}
