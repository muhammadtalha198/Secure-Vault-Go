package domain

import (
	"time"

	"github.com/google/uuid"
)

// Document stores metadata only — encrypted file lives in S3
type Document struct {
	ID               uuid.UUID `json:"id"`
	OwnerID          uuid.UUID `json:"owner_id"`
	OriginalFilename string    `json:"original_filename"`
	FileSize         int64     `json:"file_size"`
	FileMimeType     string    `json:"file_mime_type"`
	FilePath         string    `json:"file_path"` // Internal, don't expose
	CreatedAt        time.Time `json:"created_at"`
}
