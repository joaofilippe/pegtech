package entities

import (
	"errors"
	"time"
)

var (
	ErrLockerNotAvailable = errors.New("locker is not available")
	ErrLockerNotReserved  = errors.New("locker is not reserved")
)

type LockerStatus string

const (
	LockerStatusAvailable   LockerStatus = "AVAILABLE"
	LockerStatusOccupied    LockerStatus = "OCCUPIED"
	LockerStatusReserved    LockerStatus = "RESERVED"
	LockerStatusMaintenance LockerStatus = "MAINTENANCE"
)

type Locker struct {
	ID            int
	Number        string
	Location      string
	Status        LockerStatus
	Package       *Package
	Client        *User
	ReservedAt    *time.Time
	OccupiedAt    *time.Time
	OccupiedUntil *time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func NewLocker(id int,number, size, location string) *Locker {
	return &Locker{
		ID:        id,
		Number:    number,
		Location:  location,
		Status:    LockerStatusAvailable,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (l *Locker) Reserve(client *User) error {
	if l.Status != LockerStatusAvailable {
		return ErrLockerNotAvailable
	}

	now := time.Now()
	l.Status = LockerStatusReserved
	l.Client = client
	l.ReservedAt = &now
	l.UpdatedAt = now

	return nil
}

func (l *Locker) Occupy(package_ *Package) error {
	if l.Status != LockerStatusReserved {
		return ErrLockerNotReserved
	}

	now := time.Now()
	l.Status = LockerStatusOccupied
	l.Package = package_
	l.OccupiedAt = &now
	l.UpdatedAt = now

	return nil
}

func (l *Locker) Release() {
	l.Status = LockerStatusAvailable
	l.Package = nil
	l.Client = nil
	l.ReservedAt = nil
	l.OccupiedAt = nil
	l.UpdatedAt = time.Now()
}

func (l *Locker) SetMaintenance() {
	l.Status = LockerStatusMaintenance
	l.Package = nil
	l.Client = nil
	l.ReservedAt = nil
	l.OccupiedAt = nil
	l.UpdatedAt = time.Now()
}

func (l *Locker) SetAvailable() {
	l.Status = LockerStatusAvailable
	l.UpdatedAt = time.Now()
}
