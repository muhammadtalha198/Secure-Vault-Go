package domain

import (
	"time"

	"github.com/google/uuid"
)

// VaultIdentity stores cryptographic identity
// Server knows public key, never knows private key
type VaultIdentity struct {
	ID                        uuid.UUID `json:"id"`
	UserID                    uuid.UUID `json:"user_id"`
	PublicKey                 string    `json:"public_key"`
	EncryptedPrivateKeyBackup string    `json:"-"` // Never serialize
	KeyAlgorithm              string    `json:"key_algorithm"`
	CreatedAt                 time.Time `json:"created_at"`
}
