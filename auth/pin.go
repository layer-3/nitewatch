package auth

import (
	"crypto/rand"
	"fmt"
	"hash/crc64"
	"math/big"
	"time"
)

const (
	PINLength     = 6
	PINExpiration = 10 * time.Minute
)

// pinManager handles PIN generation and validation
type pinManager struct {
	expirationDuration time.Duration
}

// newPINManager creates a new PIN manager instance
func newPINManager() *pinManager {
	return &pinManager{
		expirationDuration: PINExpiration,
	}
}

// generatePIN generates a random 6-digit PIN
func (p *pinManager) generatePIN() (string, error) {
	// Generate a random number between 0 and 999999
	max := big.NewInt(1000000)
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", err
	}

	// Format as 6-digit string with leading zeros
	pin := fmt.Sprintf("%06d", n.Int64())
	return pin, nil
}

// hashPIN creates a secure hash of the PIN using CRC64
func (p *pinManager) hashPIN(pin string) uint64 {
	table := crc64.MakeTable(crc64.ISO)
	return crc64.Checksum([]byte(pin), table)
}

// validatePIN validates a PIN against its hash
func (p *pinManager) validatePIN(pin string, hashedPIN uint64) bool {
	return p.hashPIN(pin) == hashedPIN
}

// isExpired checks if a PIN has expired based on creation time
func (p *pinManager) isExpired(createdAt time.Time) bool {
	return time.Since(createdAt) > p.expirationDuration
}

// validatePINWithExpiry validates both the PIN and its expiration
func (p *pinManager) validatePINWithExpiry(pin string, hashedPIN uint64, createdAt time.Time) error {
	if p.isExpired(createdAt) {
		return ErrPINExpired
	}

	if !p.validatePIN(pin, hashedPIN) {
		return ErrInvalidPIN
	}

	return nil
}
