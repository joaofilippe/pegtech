package lockerusecases

import (
	"github.com/joaofilippe/pegtech/internal/domain/irepositories"
)

// ReleaseLockerCase handles locker release
type ReleaseLockerCase struct {
	lockerRepo irepositories.LockerRepository
}

// NewReleaseLockerCase creates a new instance of ReleaseLockerCase
func NewReleaseLockerCase(lockerRepo irepositories.LockerRepository) *ReleaseLockerCase {
	return &ReleaseLockerCase{
		lockerRepo: lockerRepo,
	}
}

// Execute performs the locker release operation
func (uc *ReleaseLockerCase) Execute(lockerID int, packageCode string) error {
	return uc.lockerRepo.ReleaseLocker(lockerID, packageCode)
}
