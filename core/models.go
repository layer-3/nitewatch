package core

import (
	"time"

	"github.com/google/uuid"
)

// User represents a user in the system
type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	CreatedAt time.Time

	PinSentAt  *time.Time
	LastSeenAt *time.Time
	NotifiedAt *time.Time

	Email    string `gorm:"uniqueIndex;not null"`
	UserPin  *int64
	Metadata []byte `gorm:"type:jsonb"`
}

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
