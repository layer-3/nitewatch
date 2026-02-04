package auth

import (
	"context"
	"errors"
	"time"

	"github.com/layer-3/nitewatch/core/log"

	"github.com/google/uuid"
	"github.com/layer-3/nitewatch/config"
	"github.com/layer-3/nitewatch/core"
	"gorm.io/gorm"
)

// Service provides authentication functionality
type Service struct {
	db            *gorm.DB
	jwtManager    *jwtManager
	pinManager    *pinManager
	emailService  Mailer
	deviceManager *deviceManager
}

// NewService creates a new authentication service
func NewService(db *gorm.DB, authConfig *config.Auth) *Service {
	if authConfig.JWTSecret == "" {
		log.Fatal().Msg("AUTH_JWT_SECRET environment variable is required")
	}

	// Choose mailer based on ResendAPIKey availability
	var mailer Mailer
	if authConfig.ResendAPIKey == "" {
		log.Println("WARNING: ResendAPIKey not configured, using console mailer for development")
		mailer = newConsoleMailer()
	} else {
		mailer = newResendMailer(authConfig.ResendAPIKey, authConfig.FromEmail)
	}

	return &Service{
		db:            db,
		jwtManager:    newJWTManager(authConfig.JWTSecret, authConfig.AccessDuration, authConfig.RefreshDuration),
		pinManager:    newPINManager(),
		emailService:  mailer,
		deviceManager: newDeviceManager(db),
	}
}

// LoginRequest represents the login request structure
type LoginRequest struct {
	Email string `json:"email" binding:"required,email"`
}

// LoginResponse represents the login response structure
type LoginResponse struct {
	Message   string `json:"message"`
	ExpiresIn int    `json:"expires_in"`
}

// SendLoginPIN generates and sends a login PIN to the user
func (s *Service) SendLoginPIN(email string, userAgent string) (*LoginResponse, error) {
	// Generate PIN
	pin, err := s.pinManager.generatePIN()
	if err != nil {
		return nil, err
	}

	// Hash the PIN
	hashedPIN := s.pinManager.hashPIN(pin)
	pinInt := int64(hashedPIN)
	now := time.Now()

	// Check if user exists, create if not
	var user core.User
	err = s.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Create new user
			user = core.User{
				Email:     email,
				UserPin:   &pinInt,
				PinSentAt: &now,
			}
			if err := s.db.Create(&user).Error; err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	} else {
		// Update existing user with new PIN
		if err := s.db.Model(&user).Updates(map[string]interface{}{
			"user_pin":    &pinInt,
			"pin_sent_at": &now,
		}).Error; err != nil {
			return nil, err
		}
	}

	// Send PIN via email
	if err := s.emailService.Send(context.Background(), map[string]any{
		"email": email,
		"pin":   pin,
	}); err != nil {
		return nil, err
	}

	// Update user's last seen
	s.db.Model(&user).Update("last_seen_at", &now)

	return &LoginResponse{
		Message:   "PIN sent to email",
		ExpiresIn: int(s.pinManager.expirationDuration.Seconds()),
	}, nil
}

// VerifyPINRequest represents the PIN verification request
type VerifyPINRequest struct {
	Email string `json:"email" binding:"required,email"`
	PIN   string `json:"pin" binding:"required,len=6"`
}

// VerifyPINResponse represents the PIN verification response
type VerifyPINResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// VerifyPIN verifies the PIN and returns JWT tokens
func (s *Service) VerifyPIN(email string, pin string, userAgent string) (*VerifyPINResponse, string, error) {
	// Get user
	var user core.User
	err := s.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, "", err
	}

	// Validate PIN with expiry
	if user.UserPin == nil || user.PinSentAt == nil {
		return nil, "", errors.New("no PIN found for user")
	}
	err = s.pinManager.validatePINWithExpiry(pin, uint64(*user.UserPin), *user.PinSentAt)
	if err != nil {
		return nil, "", err
	}

	// Create a new device for the user
	device, err := s.deviceManager.createDevice(
		&user.ID,
		userAgent,
		uuid.New(),
		time.Now().Add(s.jwtManager.refreshDuration),
	)
	if err != nil {
		return nil, "", err
	}

	// Generate tokens
	accessToken, refreshToken, err := s.jwtManager.generateTokenPair(email, device.Token)
	if err != nil {
		return nil, "", err
	}

	// Clear the PIN from the user record
	if err := s.db.Model(&user).Updates(map[string]interface{}{
		"user_pin":    nil,
		"pin_sent_at": nil,
	}).Error; err != nil {
		// Log error but don't fail the login
		log.Printf("Failed to clear PIN for user %s: %v", email, err)
	}

	// Update user's last seen
	now := time.Now()
	s.db.Model(&user).Update("last_seen_at", &now)

	return &VerifyPINResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, refreshToken, nil
}

// RefreshTokens generates new tokens using a refresh token
func (s *Service) RefreshTokens(refreshToken string) (*VerifyPINResponse, string, error) {
	// Validate refresh token
	claims, err := s.jwtManager.validateToken(refreshToken, RefreshTokenType)
	if err != nil {
		return nil, "", err
	}

	// Validate device
	deviceID, err := uuid.Parse(claims.Subject)
	if err != nil {
		return nil, "", err
	}
	device, err := s.deviceManager.validateDevice(deviceID)
	if err != nil {
		return nil, "", err
	}

	// Generate new tokens with same UUID for refresh token
	accessToken, newRefreshToken, err := s.jwtManager.refreshTokens(claims.Email, deviceID)
	if err != nil {
		return nil, "", err
	}

	// Update device expiration
	expiresAt := time.Now().Add(s.jwtManager.refreshDuration)
	if err := s.deviceManager.updateDeviceToken(device.ID, deviceID, expiresAt); err != nil {
		return nil, "", err
	}

	// Update user's last seen
	var user core.User
	if device.UserID != nil {
		if err := s.db.First(&user, *device.UserID).Error; err == nil {
			now := time.Now()
			s.db.Model(&user).Update("last_seen_at", &now)
		}
	}

	return &VerifyPINResponse{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
	}, newRefreshToken, nil
}

// Logout revokes a specific device or all devices
func (s *Service) Logout(refreshToken string, revokeAll bool) error {
	// Extract sub from refresh token
	claims, err := s.jwtManager.validateToken(refreshToken, RefreshTokenType)
	if err != nil {
		return err
	}

	// Parse token UUID
	deviceID, err := uuid.Parse(claims.Subject)
	if err != nil {
		return err
	}

	// Get device to find user
	device, err := s.deviceManager.validateDevice(deviceID)
	if err != nil {
		// If device not found or already revoked, consider it a success
		if errors.Is(err, ErrDeviceNotFound) || errors.Is(err, ErrTokenRevoked) {
			return nil
		}
		return err
	}

	if revokeAll {
		// Revoke all user devices
		return s.deviceManager.revokeAllUserDevices(device.UserID)
	} else {
		// Revoke only this device
		return s.deviceManager.revokeDevice(deviceID)
	}
}

// ValidateAccessToken validates an access token and returns the user email
func (s *Service) ValidateAccessToken(token string) (string, error) {
	claims, err := s.jwtManager.validateToken(token, AccessTokenType)
	if err != nil {
		return "", err
	}
	return claims.Email, nil
}
