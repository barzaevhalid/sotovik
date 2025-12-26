package auth

import "errors"

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrForbidden          = errors.New("forbidden")
	ErrUserBlocked        = errors.New("user is blocked")
)
