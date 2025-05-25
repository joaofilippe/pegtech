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
func (uc *RegisterLockerCase) Execute(lockerID int, ports []int) error {
	for _, port := range ports {
		locker := &entities.Port{
			Locker: lockerID,
			Port:   port,
			Status: entities.LockerStatusAvailable,
		}
		err := uc.lockerRepo.SaveLocker(locker)
		if err != nil {
			return err
		}
	}

	return nil
}
