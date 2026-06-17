package domain

import (
	"time"

	"github.com/google/uuid"
)

// User represents an application account
// Separate from vault cryptographic identity
type User struct {
	ID              uuid.UUID `json:"id"`
	Email           string    `json:"email"`
	PasswordHash    string    `json:"-"` // Never serialize to JSON
	IsEmailVerified bool      `json:"is_email_verified,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// Session tracks issued refresh tokens
type Session struct {
	ID               uuid.UUID `json:"id"`
	UserID           uuid.UUID `json:"user_id"`
	RefreshTokenHash string    `json:"-"` // Never serialize
	DeviceInfo       string    `json:"device_info"`
	IsRevoked        bool      `json:"is_revoked"`
	ExpiresAt        time.Time `json:"expires_at"`
	LastUsedAt       time.Time `json:"last_used_at"`
}
