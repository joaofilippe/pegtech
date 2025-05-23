package errors

import "errors"

var (
	ErrFoundNoLockers        = errors.New("no lockers found")
	ErrLockerAlreadyOccupied = errors.New("locker is already occupied")
	ErrNoAvailableLockers    = errors.New("no available lockers found")
)
