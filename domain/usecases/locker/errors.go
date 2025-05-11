package lockerusecases

import "errors"

// Common errors for locker operations
var (
	ErrNoAvailableLockers = errors.New("no available lockers")
	ErrLockerNotFound     = errors.New("locker not found")
	ErrInvalidPassword    = errors.New("invalid password")
)
