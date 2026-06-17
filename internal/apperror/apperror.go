package apperrors

import (
	"errors"
	"fmt"
)

// Sentinel errors (simple, static)
var (
	ErrVaultNotFound      = errors.New("vault not found")
	ErrDocumentNotFound   = errors.New("document not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrSeedPhraseRequired = errors.New("seed phrase required for recovery")
)

// AppError - structured error with HTTP mapping
type AppError struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	StatusCode int    `json:"-"`
	Err        error  `json:"-"`
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

// Predefined error constructors
func VaultNotFound(id string) *AppError {
	return &AppError{
		Code:       "VAULT_NOT_FOUND",
		Message:    fmt.Sprintf("Vault %s not found", id),
		StatusCode: 404,
		Err:        ErrVaultNotFound,
	}
}

func DocumentNotFound(id string) *AppError {
	return &AppError{
		Code:       "DOCUMENT_NOT_FOUND",
		Message:    fmt.Sprintf("Document %s not found", id),
		StatusCode: 404,
		Err:        ErrDocumentNotFound,
	}
}

func Unauthorized(msg string) *AppError {
	return &AppError{
		Code:       "UNAUTHORIZED",
		Message:    msg,
		StatusCode: 401,
	}
}

func BadRequest(msg string) *AppError {
	return &AppError{
		Code:       "BAD_REQUEST",
		Message:    msg,
		StatusCode: 400,
	}
}

func Internal(msg string, err error) *AppError {
	return &AppError{
		Code:       "INTERNAL_ERROR",
		Message:    msg,
		StatusCode: 500,
		Err:        err,
	}
}
