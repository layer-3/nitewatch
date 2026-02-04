package auth

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/layer-3/nitewatch/core"
	"gorm.io/gorm"
)

// deviceManager handles device and session management
type deviceManager struct {
	db *gorm.DB
}

// newDeviceManager creates a new device manager instance
func newDeviceManager(db *gorm.DB) *deviceManager {
	return &deviceManager{
		db: db,
	}
}

// createDevice creates a new device record for authentication
func (d *deviceManager) createDevice(userID *uuid.UUID, userAgent string, token uuid.UUID, expiresAt time.Time) (*core.Device, error) {
	device := &core.Device{
		UserID:    userID,
		UserAgent: userAgent,
		Token:     token,
		ExpiresAt: expiresAt,
	}

	if err := d.db.Create(device).Error; err != nil {
		return nil, err
	}

	return device, nil
}

// validateDevice checks if a device/token is valid and not revoked
func (d *deviceManager) validateDevice(token uuid.UUID) (*core.Device, error) {
	var device core.Device
	err := d.db.Where("token = ?", token).First(&device).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrDeviceNotFound
		}
		return nil, err
	}

	// Check if token is revoked
	if device.RevokedAt != nil {
		return nil, ErrTokenRevoked
	}

	// Check if token is expired
	if time.Now().After(device.ExpiresAt) {
		return nil, ErrTokenExpired
	}

	return &device, nil
}

// updateDeviceToken updates the token and expiration for a device
func (d *deviceManager) updateDeviceToken(deviceID uuid.UUID, newToken uuid.UUID, newExpiresAt time.Time) error {
	return d.db.Model(&core.Device{}).
		Where("id = ?", deviceID).
		Updates(map[string]interface{}{
			"token":      newToken,
			"expires_at": newExpiresAt,
		}).Error
}

// revokeDevice revokes a specific device/token
func (d *deviceManager) revokeDevice(token uuid.UUID) error {
	now := time.Now()
	return d.db.Model(&core.Device{}).
		Where("token = ?", token).
		Update("revoked_at", &now).Error
}

// revokeAllUserDevices revokes all devices for a user
func (d *deviceManager) revokeAllUserDevices(userID *uuid.UUID) error {
	now := time.Now()
	return d.db.Model(&core.Device{}).
		Where("user_id = ? AND revoked_at IS NULL", userID).
		Update("revoked_at", &now).Error
}
