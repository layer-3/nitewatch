package core

import (
	"time"

	"github.com/google/uuid"
)

// Device represents a user's device/session
type Device struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID    *uuid.UUID `gorm:"type:uuid"`
	CreatedAt time.Time
	ExpiresAt time.Time
	RevokedAt *time.Time
	UserAgent string
	Token     uuid.UUID `gorm:"type:uuid;index"`
}
