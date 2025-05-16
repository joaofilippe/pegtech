package lockerusecases

import (
	"github.com/joaofilippe/pegtech/internal/domain/entities"
	"github.com/joaofilippe/pegtech/internal/domain/irepositories"
)

// UpdateLockerStatusCase handles locker status updates
type UpdateLockerStatusCase struct {
	lockerRepo irepositories.LockerRepository
}

// NewUpdateLockerStatusCase creates a new instance of UpdateLockerStatusCase
func NewUpdateLockerStatusCase(lockerRepo irepositories.LockerRepository) *UpdateLockerStatusCase {
	return &UpdateLockerStatusCase{
		lockerRepo: lockerRepo,
	}
}

// Execute performs the locker status update operation
func (uc *UpdateLockerStatusCase) Execute(id string, status entities.LockerStatus) error {
	// Check if locker exists
	_, err := uc.lockerRepo.GetLocker(id)
	if err != nil {
		return ErrLockerNotFound
	}

	return uc.lockerRepo.UpdateLockerStatus(id, status)
}
