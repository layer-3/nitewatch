package auth

import (
	"strings"
	"time"

	"github.com/layer-3/nitewatch/core/log"

	"github.com/google/uuid"
)

// Session represents a user's authenticated session
type Session struct {
	Email     string
	Device    uuid.UUID
	ExpiresAt time.Time
}

// GetSession validates an access token and returns the session details
func (s *Service) GetSession(accessToken string) (*Session, error) {
	// Validate the access token
	claims, err := s.jwtManager.validateToken(accessToken, AccessTokenType)
	if err != nil {
		return nil, err
	}

	// Parse the device token from the subject claim
	deviceToken, err := uuid.Parse(claims.Subject)
	if err != nil {
		return nil, err
	}

	// Create the session struct
	session := &Session{
		Email:     claims.Email,
		Device:    deviceToken,
		ExpiresAt: claims.ExpiresAt.Time,
	}

	return session, nil
}

// ParseSession parses the Authorization header and returns a Session if valid, otherwise nil.
func (s *Service) ParseSession(authHeader string) *Session {
	if authHeader == "" {
		return nil
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return nil
	}

	session, err := s.GetSession(parts[1])
	if err != nil {
		log.Printf("Failed to validate access token: %v", err)
		return nil
	}

	return session
}
