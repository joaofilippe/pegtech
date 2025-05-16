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
	ID           uuid.UUID
	TrackingCode string
	Description  string
	Weight       float64
	Dimensions   struct {
		Length float64
		Width  float64
		Height float64
	}
	Status          PackageStatus
	Sender          *Client
	Recipient       *Client
	Locker          *Locker
	PickupPassword  string
	PickupExpiresAt time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func NewPackage(description string, weight float64, length, width, height float64, sender, recipient *Client) *Package {
	return &Package{
		ID:           uuid.New(),
		TrackingCode: generateTrackingCode(),
		Description:  description,
		Weight:       weight,
		Dimensions: struct {
			Length float64
			Width  float64
			Height float64
		}{
			Length: length,
			Width:  width,
			Height: height,
		},
		Status:    PackageStatusPending,
		Sender:    sender,
		Recipient: recipient,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (p *Package) UpdateStatus(status PackageStatus) {
	p.Status = status
	p.UpdatedAt = time.Now()
}

func (p *Package) UpdateDimensions(length, width, height float64) {
	p.Dimensions.Length = length
	p.Dimensions.Width = width
	p.Dimensions.Height = height
	p.UpdatedAt = time.Now()
}

func (p *Package) UpdateWeight(weight float64) {
	p.Weight = weight
	p.UpdatedAt = time.Now()
}

func generateTrackingCode() string {
	return uuid.New().String()[:8]
}
