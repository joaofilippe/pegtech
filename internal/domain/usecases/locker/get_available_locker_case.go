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
	return c.lockerRepo.GetAvailableLockers()
}
