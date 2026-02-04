package config

import (
	"fmt"

	"github.com/layer-3/nitewatch/core/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database holds database configuration
type Database struct {
	Host     string `env:"DB_HOST"`
	Port     string `env:"DB_PORT" env-default:"5432"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Name     string `env:"DB_NAME"`
	SSLMode  string `env:"DB_SSLMODE" env-default:"disable"`
	TimeZone string `env:"DB_TIMEZONE" env-default:"UTC"`
}

// GetDSN returns the database connection string based on the configuration
func (c *Database) GetDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		c.Host, c.Port, c.User, c.Password, c.Name, c.SSLMode, c.TimeZone)
}

// GetGormDialector returns the appropriate GORM dialector based on the database type
func (c *Database) GetGormDialector() gorm.Dialector {
	return postgres.Open(c.GetDSN())
}

// InitDBWithConfig initializes a connection to the database using the provided config
func InitDBWithConfig(cfg *Database, lvl log.LogLevel) (*gorm.DB, error) {
	gormLogger := log.NewGormLogger(lvl)

	// Set GORM configuration
	gormConfig := &gorm.Config{
		Logger: gormLogger,
	}

	// Connect to the database
	db, err := gorm.Open(cfg.GetGormDialector(), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	return db, nil
}
