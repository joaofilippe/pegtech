package lockerusecases

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"time"

	"github.com/joaofilippe/pegtech/domain/entities"
	"github.com/joaofilippe/pegtech/domain/irepositories"
)

var (
	ErrNoAvailableLockers = errors.New("no available lockers")
	ErrLockerNotFound     = errors.New("locker not found")
	ErrInvalidPassword    = errors.New("invalid password")
)

type LockerUseCase struct {
	lockerRepo  irepositories.LockerRepository
	packageRepo irepositories.PackageRepository
}

func NewLockerUseCase(lockerRepo irepositories.LockerRepository, packageRepo irepositories.PackageRepository) *LockerUseCase {
	return &LockerUseCase{
		lockerRepo:  lockerRepo,
		packageRepo: packageRepo,
	}
}

func (uc *LockerUseCase) RegisterLocker(id string, size string) error {
	locker := &entities.Locker{
		ID:     id,
		Status: entities.LockerAvailable,
		Size:   size,
	}
	return uc.lockerRepo.SaveLocker(locker)
}

func (uc *LockerUseCase) GetAvailableLocker(size string) (*entities.Locker, error) {
	return uc.lockerRepo.GetAvailableLocker(size)
}

func (uc *LockerUseCase) RegisterPackage(trackingCode string, size string) (*entities.Package, error) {
	locker, err := uc.GetAvailableLocker(size)
	if err != nil {
		return nil, err
	}

	password := generatePassword()
	expiresAt := time.Now().Add(24 * time.Hour)

	pkg := &entities.Package{
		ID:           generateID(),
		TrackingCode: trackingCode,
		LockerID:     locker.ID,
		Password:     password,
		CreatedAt:    time.Now(),
		ExpiresAt:    expiresAt,
		Status:       "registered",
	}

	if err := uc.packageRepo.SavePackage(pkg); err != nil {
		return nil, err
	}

	if err := uc.lockerRepo.UpdateLockerStatus(locker.ID, entities.LockerOccupied); err != nil {
		return nil, err
	}

	return pkg, nil
}

func (uc *LockerUseCase) GetPackagePickupInfo(trackingCode string) (*entities.PackagePickup, error) {
	pkg, err := uc.packageRepo.GetPackageByTrackingCode(trackingCode)
	if err != nil {
		return nil, err
	}

	return &entities.PackagePickup{
		LockerID:  pkg.LockerID,
		Password:  pkg.Password,
		ExpiresAt: pkg.ExpiresAt,
	}, nil
}

func (uc *LockerUseCase) OpenLocker(lockerID string, password string) error {
	_, err := uc.lockerRepo.GetLocker(lockerID)
	if err != nil {
		return ErrLockerNotFound
	}

	pkg, err := uc.packageRepo.GetPackageByTrackingCode(lockerID)
	if err != nil {
		return ErrInvalidPassword
	}

	if pkg.Password != password {
		return ErrInvalidPassword
	}

	if err := uc.lockerRepo.UpdateLockerStatus(lockerID, entities.LockerAvailable); err != nil {
		return err
	}

	return uc.packageRepo.DeletePackage(pkg.ID)
}

func (uc *LockerUseCase) ListLockers() ([]*entities.Locker, error) {
	return uc.lockerRepo.ListLockers()
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
