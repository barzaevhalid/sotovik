package domain

import "errors"

var (
	ErrNotFound      = errors.New("not found")
	ErrForbidden     = errors.New("forbidden")
	ErrConflict      = errors.New("conflict")
	ErrInvalidInuput = errors.New("invalid input")
)

var (
	// User errors
	ErrUserNotFound       = errors.New("user not found")
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserDisabled       = errors.New("user account is disabled")

	// Database errors
	ErrDatabase = errors.New("database error")

	// Validation errors
	ErrInvalidInput = errors.New("invalid input")
)
