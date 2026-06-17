package dtos

import (
	"time"

	"github.com/google/uuid"
)

// RegisterRequest for user signup
type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=12,max=128"`
}

// LoginRequest for user authentication
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// CreateVaultRequest for setting up cryptographic identity
type CreateVaultRequest struct {
	PublicKey                 string `json:"public_key" validate:"required,base64"`
	EncryptedPrivateKeyBackup string `json:"encrypted_private_key_backup" validate:"required"`
}

// VerifyChallengeRequest for proving vault ownership
type VerifyChallengeRequest struct {
	Challenge string `json:"challenge" validate:"required"`
	Signature string `json:"signature" validate:"required,base64"`
}

// UploadMeta for document upload metadata
type UploadMeta struct {
	OriginalFilename string `json:"original_filename" validate:"required,max=512"`
	FileMimeType     string `json:"file_mime_type" validate:"required"`
	EncryptedAesKey  string `json:"encrypted_aes_key" validate:"required,base64"`
}

// ShareRequest for sharing a document
type ShareRequest struct {
	RecipientEmail              string     `json:"recipient_email" validate:"required,email"`
	EncryptedAesKeyForRecipient string     `json:"encrypted_aes_key_for_recipient" validate:"required,base64"`
	Permission                  string     `json:"permission" validate:"required,oneof=view download"`
	ExpiresAt                   *time.Time `json:"expires_at,omitempty" validate:"omitempty,gtfield=CreatedAt"`
}

// CreateDocumentParams for repository layer
type CreateDocumentParams struct {
	ID               string
	OwnerID          string
	OriginalFilename string
	FileSize         int64
	FileMimeType     string
	StorageKey       string
	EncryptedAesKey  string
}

// LoginResponse returned after successful authentication
type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"` // seconds
}

// DocumentResponse for API responses (no sensitive fields)
type DocumentResponse struct {
	ID               uuid.UUID `json:"id"`
	OriginalFilename string    `json:"original_filename"`
	FileSize         int64     `json:"file_size"`
	FileMimeType     string    `json:"file_mime_type"`
	CreatedAt        time.Time `json:"created_at"`
}

// UserResponse for profile endpoints
type UserResponse struct {
	ID               uuid.UUID `json:"id"`
	Email            string    `json:"email"`
	SubscriptionTier string    `json:"subscription_tier"`
	IsEmailVerified  bool      `json:"is_email_verified"`
}

// StorageUsageResponse for quota information
type StorageUsageResponse struct {
	UsedGB  int64 `json:"used_gb"`
	QuotaGB int64 `json:"quota_gb"`
}

// ShareResponse for listing shares
type ShareResponse struct {
	ID         uuid.UUID  `json:"id"`
	SharedWith string     `json:"shared_with_email"`
	Permission string     `json:"permission"`
	ExpiresAt  *time.Time `json:"expires_at,omitempty"`
	CreatedAt  time.Time  `json:"created_at"`
}
