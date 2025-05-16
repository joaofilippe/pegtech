package lockerusecases

import (
	"github.com/google/uuid"
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
func (uc *RegisterLockerCase) Execute(id string, size string) error {
	lockerID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	locker := &entities.Locker{
		ID:     lockerID,
		Status: entities.LockerStatusAvailable,
		Size:   size,
	}
	return uc.lockerRepo.SaveLocker(locker)
}
