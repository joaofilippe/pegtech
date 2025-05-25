package lockerusecases

import "errors"

// Common errors for locker operations
var (
	ErrNoAvailableLockers    = errors.New("no available lockers")
	ErrNoAvailablePorts      = errors.New("no available ports")
	ErrLockerNotFound        = errors.New("locker not found")
	ErrInvalidPassword       = errors.New("invalid password")
	ErrLockerAlreadyOccupied = errors.New("locker already occupied")
	ErrAllLockersOccupied    = errors.New("all lockers are occupied")
	ErrFoundNoLockers        = errors.New("no lockers found")
)
