package services

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"sync"
	"time"

	"github.com/joaofilippe/pegtech/application/models"
)

var (
	ErrNoAvailableLockers = errors.New("no available lockers")
	ErrLockerNotFound     = errors.New("locker not found")
	ErrInvalidPassword    = errors.New("invalid password")
)

type LockerService struct {
	lockers  map[string]*models.Locker
	packages map[string]*models.Package
	mu       sync.RWMutex
}

func NewLockerService() *LockerService {
	return &LockerService{
		lockers:  make(map[string]*models.Locker),
		packages: make(map[string]*models.Package),
	}
}

func (s *LockerService) RegisterLocker(id string, size string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.lockers[id] = &models.Locker{
		ID:     id,
		Status: models.LockerAvailable,
		Size:   size,
	}
}

func (s *LockerService) GetAvailableLocker(size string) (*models.Locker, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, locker := range s.lockers {
		if locker.Status == models.LockerAvailable && locker.Size == size {
			return locker, nil
		}
	}
	return nil, ErrNoAvailableLockers
}

func (s *LockerService) RegisterPackage(reg models.PackageRegistration) (*models.Package, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	locker, err := s.GetAvailableLocker(reg.Size)
	if err != nil {
		return nil, err
	}

	password := generatePassword()
	expiresAt := time.Now().Add(24 * time.Hour)

	pkg := &models.Package{
		ID:           generateID(),
		TrackingCode: reg.TrackingCode,
		LockerID:     locker.ID,
		Password:     password,
		CreatedAt:    time.Now(),
		ExpiresAt:    expiresAt,
		Status:       "registered",
	}

	s.packages[pkg.ID] = pkg
	s.lockers[locker.ID].Status = models.LockerOccupied

	return pkg, nil
}

func (s *LockerService) GetPackagePickupInfo(trackingCode string) (*models.PackagePickup, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, pkg := range s.packages {
		if pkg.TrackingCode == trackingCode {
			return &models.PackagePickup{
				LockerID:  pkg.LockerID,
				Password:  pkg.Password,
				ExpiresAt: pkg.ExpiresAt,
			}, nil
		}
	}
	return nil, errors.New("package not found")
}

func (s *LockerService) OpenLocker(lockerID, password string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	locker, exists := s.lockers[lockerID]
	if !exists {
		return ErrLockerNotFound
	}

	for _, pkg := range s.packages {
		if pkg.LockerID == lockerID && pkg.Password == password {
			locker.Status = models.LockerAvailable
			delete(s.packages, pkg.ID)
			return nil
		}
	}

	return ErrInvalidPassword
}

func generatePassword() string {
	bytes := make([]byte, 4)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func generateID() string {
	bytes := make([]byte, 8)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}
