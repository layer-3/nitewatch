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
