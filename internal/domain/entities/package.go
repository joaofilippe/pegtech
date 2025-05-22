package entities

import (
	"time"

	"github.com/google/uuid"
)

type PackageStatus string

const (
	PackageStatusPending   PackageStatus = "PENDING"
	PackageStatusInTransit PackageStatus = "IN_TRANSIT"
	PackageStatusDelivered PackageStatus = "DELIVERED"
	PackageStatusReturned  PackageStatus = "RETURNED"
)

type Package struct {
	ID              uuid.UUID
	TrackingCode    string
	Description     string
	Status          PackageStatus
	Recipient       *User
	Locker          *Locker
	PickupPassword  string
	PickupExpiresAt time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func NewPackage(description string, recipient *User) *Package {
	return &Package{
		ID:           uuid.New(),
		TrackingCode: generateTrackingCode(),
		Description:  description,
		Status:       PackageStatusPending,
		Recipient:    recipient,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

func (p *Package) UpdateStatus(status PackageStatus) {
	p.Status = status
	p.UpdatedAt = time.Now()
}

func generateTrackingCode() string {
	return uuid.New().String()[:8]
}
