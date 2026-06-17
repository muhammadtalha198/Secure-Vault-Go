package apperror

import (
	"errors"
	"fmt"
	"net/http"
)

type AppError struct {
	StatusCode int
	Message    string
	Err        error
}

// Error implements the error interface
func (e *AppError) Error() string {
	return fmt.Sprintf("code: %s, message: %s, error: %v", e.StatusCode, e.Message, e.Err)
}

// Helper function to craete specific error
func ErrnotFound(msg string) *AppError {
	return &AppError{
		StatusCode: http.StatusNotFound,
		Message:    msg,
		Err:        errors.New("not found"),
	}
}

func ErrUnauthorized(msg string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		Message:    msg,
		Err:        errors.New("unauthorized"),
	}
}

func ErrTooManyRequests(msg string) *AppError {
	return &AppError{
		StatusCode: http.StatusTooManyRequests,
		Message:    msg,
		Err:        errors.New("too many requests"),
	}
}
