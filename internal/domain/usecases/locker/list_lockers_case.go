package lockerusecases

import (
	"github.com/joaofilippe/pegtech/domain/entities"
	"github.com/joaofilippe/pegtech/domain/irepositories"
)

// ListLockersCase handles listing all lockers
type ListLockersCase struct {
	lockerRepo irepositories.LockerRepository
}

// NewListLockersCase creates a new instance of ListLockersCase
func NewListLockersCase(lockerRepo irepositories.LockerRepository) *ListLockersCase {
	return &ListLockersCase{
		lockerRepo: lockerRepo,
	}
}

// Execute performs the locker listing operation
func (uc *ListLockersCase) Execute() ([]*entities.Locker, error) {
	return uc.lockerRepo.ListLockers()
}
