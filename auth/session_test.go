package auth

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
)

// newTestService is a helper to create a Service with a configured jwtManager for testing.
func newTestService(secret string) *Service {
	return &Service{
		jwtManager: newJWTManager(secret, time.Minute*15, time.Hour*24),
	}
}

func TestParseSession(t *testing.T) {
	service := newTestService("test-secret")
	email := "test@example.com"
	deviceID := uuid.New()

	t.Run("valid session", func(t *testing.T) {
		accessToken, _, err := service.jwtManager.generateToken(email, deviceID, AccessTokenType, time.Minute*5)
		if err != nil {
			t.Fatalf("unexpected error generating token: %v", err)
		}

		header := fmt.Sprintf("Bearer %s", accessToken)
		session := service.ParseSession(header)

		if session == nil {
			t.Fatal("session should not be nil")
		}
		if session.Email != email {
			t.Errorf("expected email %s, got %s", email, session.Email)
		}
		if session.Device != deviceID {
			t.Errorf("expected device ID %s, got %s", deviceID, session.Device)
		}
	})

	t.Run("empty header", func(t *testing.T) {
		session := service.ParseSession("")
		if session != nil {
			t.Fatalf("session should be nil, got %+v", session)
		}
	})

	t.Run("malformed header no bearer", func(t *testing.T) {
		accessToken, _, err := service.jwtManager.generateToken(email, deviceID, AccessTokenType, time.Minute*5)
		if err != nil {
			t.Fatalf("unexpected error generating token: %v", err)
		}

		session := service.ParseSession(accessToken)
		if session != nil {
			t.Fatalf("session should be nil, got %+v", session)
		}
	})

	t.Run("malformed header wrong scheme", func(t *testing.T) {
		accessToken, _, err := service.jwtManager.generateToken(email, deviceID, AccessTokenType, time.Minute*5)
		if err != nil {
			t.Fatalf("unexpected error generating token: %v", err)
		}

		header := fmt.Sprintf("Basic %s", accessToken)
		session := service.ParseSession(header)
		if session != nil {
			t.Fatalf("session should be nil, got %+v", session)
		}
	})

	t.Run("expired token", func(t *testing.T) {
		expiredToken, _, err := service.jwtManager.generateToken(email, deviceID, AccessTokenType, -time.Minute*5)
		if err != nil {
			t.Fatalf("unexpected error generating token: %v", err)
		}

		header := fmt.Sprintf("Bearer %s", expiredToken)
		session := service.ParseSession(header)
		if session != nil {
			t.Fatalf("session should be nil, got %+v", session)
		}
	})

	t.Run("invalid signature", func(t *testing.T) {
		service2 := newTestService("different-secret")
		accessToken, _, err := service.jwtManager.generateToken(email, deviceID, AccessTokenType, time.Minute*5)
		if err != nil {
			t.Fatalf("unexpected error generating token: %v", err)
		}

		header := fmt.Sprintf("Bearer %s", accessToken)
		session := service2.ParseSession(header)
		if session != nil {
			t.Fatalf("session should be nil, got %+v", session)
		}
	})

	t.Run("wrong token type", func(t *testing.T) {
		refreshToken, _, err := service.jwtManager.generateToken(email, deviceID, RefreshTokenType, time.Minute*5)
		if err != nil {
			t.Fatalf("unexpected error generating token: %v", err)
		}

		header := fmt.Sprintf("Bearer %s", refreshToken)
		session := service.ParseSession(header)
		if session != nil {
			t.Fatalf("session should be nil, got %+v", session)
		}
	})

	t.Run("invalid token string", func(t *testing.T) {
		header := "Bearer invalid-token"
		session := service.ParseSession(header)
		if session != nil {
			t.Fatalf("session should be nil, got %+v", session)
		}
	})
}
