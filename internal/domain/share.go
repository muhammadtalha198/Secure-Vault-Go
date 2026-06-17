package domain

import (
	"time"

	"github.com/google/uuid"
)

// DocumentShare tracks who has access to what
type DocumentShare struct {
	ID         uuid.UUID  `json:"id"`
	DocumentID uuid.UUID  `json:"document_id"`
	SharedBy   uuid.UUID  `json:"shared_by"`
	SharedWith uuid.UUID  `json:"shared_with"`
	Permission string     `json:"permission"`
	IsRevoked  bool       `json:"is_revoked"`
	ExpiresAt  *time.Time `json:"expires_at,omitempty"`
	CreatedAt  time.Time  `json:"created_at"`
}
