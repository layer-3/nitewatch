package auth

import "errors"

var (
	// Token errors
	ErrInvalidToken     = errors.New("invalid token")
	ErrTokenExpired     = errors.New("token expired")
	ErrInvalidTokenType = errors.New("invalid token type")

	// Device errors
	ErrDeviceNotFound = errors.New("device not found")
	ErrTokenRevoked   = errors.New("token has been revoked")

	// PIN errors
	ErrInvalidPIN = errors.New("invalid PIN")
	ErrPINExpired = errors.New("PIN expired")
)
