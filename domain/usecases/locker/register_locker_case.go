package lockerusecases

import (
	"github.com/joaofilippe/pegtech/domain/entities"
	"github.com/joaofilippe/pegtech/domain/irepositories"
)

// RegisterLockerCase handles locker registration
type RegisterLockerCase struct {
	lockerRepo irepositories.LockerRepository
}

// NewRegisterLockerCase creates a new instance of RegisterLockerCase
func NewRegisterLockerCase(lockerRepo irepositories.LockerRepository) *RegisterLockerCase {
	return &RegisterLockerCase{
		lockerRepo: lockerRepo,
	}
}

// Execute performs the locker registration operation
func (uc *RegisterLockerCase) Execute(id string, size string) error {
	locker := &entities.Locker{
		ID:     id,
		Status: entities.LockerAvailable,
		Size:   size,
	}
	return uc.lockerRepo.SaveLocker(locker)
}
