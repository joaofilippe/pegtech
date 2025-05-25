package lockerusecases

import (
	"github.com/joaofilippe/pegtech/internal/domain/entities"
	"github.com/joaofilippe/pegtech/internal/domain/irepositories"
)

// UpdatePortStatusCase handles locker status updates
type UpdatePortStatusCase struct {
	lockerRepo irepositories.LockerRepository
}

// NewUpdatePortStatusCase creates a new instance of UpdateLockerStatusCase
func NewUpdatePortStatusCase(lockerRepo irepositories.LockerRepository) *UpdatePortStatusCase {
	return &UpdatePortStatusCase{
		lockerRepo: lockerRepo,
	}
}

// Execute performs the locker status update operation
func (uc *UpdatePortStatusCase) Execute(packageCode string, status entities.LockerStatus) error {
	// Check if locker exists
	locker, err := uc.lockerRepo.GetLockerByPackageCode(packageCode)
	if err != nil {
		return ErrLockerNotFound
	}

	locker.Status = status
	return uc.lockerRepo.UpdateLocker(locker)
}
