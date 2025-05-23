package lockerusecases

import (
	"github.com/joaofilippe/pegtech/internal/domain/irepositories"
)

type GetAvailableLockersCase struct {
	lockerRepo irepositories.LockerRepository
}

func NewGetAvailableLockersCase(lockerRepo irepositories.LockerRepository) *GetAvailableLockersCase {
	return &GetAvailableLockersCase{
		lockerRepo: lockerRepo,
	}
}

func (c *GetAvailableLockersCase) Execute() ([]int, error) {
	lockers, err := c.lockerRepo.ListLockers()
	if err != nil {
		return nil, err
	}

	if len(lockers) == 0 {
		return nil, ErrFoundNoLockers
	}

	return getAvailableLockerIDs(lockers)
}
