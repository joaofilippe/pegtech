package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrLockerNotAvailable = errors.New("locker is not available")
	ErrLockerNotReserved  = errors.New("locker is not reserved")
)

type LockerStatus string

const (
	LockerStatusAvailable LockerStatus = "AVAILABLE"
	LockerStatusOccupied  LockerStatus = "OCCUPIED"
	LockerStatusReserved  LockerStatus = "RESERVED"
)

type Locker struct {
	ID                     int
	Number                 string
	Location               string
	PackageCode            string
	PackagePickupPassword  string
	PackagePickupExpiresAt time.Time
	PackageUserID          uuid.UUID
	Status                 LockerStatus
	Client                 *User
	ReservedExpiration     *time.Time
	OccupiedAt             *time.Time
	OccupiedUntil          *time.Time
	CreatedAt              time.Time
	UpdatedAt              time.Time
}

func NewLocker(id int, number, size, location string) *Locker {
	now := time.Now()
	return &Locker{
		ID:                    id,
		Number:                number,
		Location:              location,
		Status:                LockerStatusAvailable,
		CreatedAt:             now,
		UpdatedAt:             now,
		ReservedExpiration:    nil,
		OccupiedAt:            nil,
		OccupiedUntil:         nil,
		PackageCode:           "",
		PackagePickupPassword: "",
		PackageUserID:         uuid.Nil,
		Client:                nil,
	}
}

func (l *Locker) Reserve(client *User, expiration *time.Duration) error {
	if l.Status != LockerStatusAvailable {
		return ErrLockerNotAvailable
	}

	now := time.Now()
	l.Status = LockerStatusReserved
	l.Client = client
	l.ReservedExpiration = &now
	if expiration != nil {
		exp := now.Add(*expiration)
		l.ReservedExpiration = &exp
	}
	l.UpdatedAt = now

	return nil
}

func (l *Locker) SetAvailable() {
	l.Status = LockerStatusAvailable
	l.UpdatedAt = time.Now()
}
