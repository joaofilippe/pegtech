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
func (uc *UpdateLockerStatusCase) Execute(id int, status entities.LockerStatus) error {
	// Check if locker exists
	locker, err := uc.lockerRepo.GetLocker(id)
	if err != nil {
		return ErrLockerNotFound
	}

	locker.Status = status
	return uc.lockerRepo.UpdateLocker(locker)
}
