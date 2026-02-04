package core

import (
	"time"

	"github.com/google/uuid"
)

// WithdrawalStatus defines the possible states of a withdrawal request
type WithdrawalStatus string

const (
	StatusCreated    WithdrawalStatus = "created"
	StatusAuthorized WithdrawalStatus = "authorized"
	StatusApproved   WithdrawalStatus = "approved"
	StatusRejected   WithdrawalStatus = "rejected"
	StatusFailed     WithdrawalStatus = "failed"
)

// Withdrawal represents a withdrawal request subject to security policy validation
type Withdrawal struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`

	// Payload Data
	UserAddress  string `gorm:"not null"`
	TokenAddress string `gorm:"not null"`
	Amount       string `gorm:"not null"` // Stored as string to preserve uint256 precision
	ChainID      int64  `gorm:"not null"`
	Nonce        int64  `gorm:"not null"`
	Email        string `gorm:"not null;index"`

	// Signatures
	UserSignature      string `gorm:"not null"`
	BrokerSignature    string `gorm:"not null"`
	NitewatchSignature *string

	// State
	Status    WithdrawalStatus `gorm:"not null;default:'created';index"`
	ErrorCode *int             `gorm:"index"`
	CreatedAt time.Time
	ExpiresAt time.Time
}
