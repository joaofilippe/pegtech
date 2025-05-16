package userusecases

import "errors"

// Common errors for user operations
var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrInvalidUserData   = errors.New("invalid user data")
	ErrUserNotFound      = errors.New("user not found")
)
