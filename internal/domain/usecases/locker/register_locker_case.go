package lockerusecases

import (
	"github.com/joaofilippe/pegtech/internal/domain/entities"
	"github.com/joaofilippe/pegtech/internal/domain/irepositories"
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
func (uc *RegisterLockerCase) Execute(id int) error {
	locker := &entities.Locker{
		ID:     id,
		Status: entities.LockerStatusAvailable,
	}
	return uc.lockerRepo.SaveLocker(locker)
}
