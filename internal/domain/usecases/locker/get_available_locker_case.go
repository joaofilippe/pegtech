package lockerusecases

import (
	"github.com/joaofilippe/pegtech/internal/domain/irepositories"
)

type GetAvailableLockerCase struct {
	lockerRepo irepositories.LockerRepository
}

func NewGetAvailableLockerCase(lockerRepo irepositories.LockerRepository) *GetAvailableLockerCase {
	return &GetAvailableLockerCase{
		lockerRepo: lockerRepo,
	}
}

func (c *GetAvailableLockerCase) Execute() ([]int, error) {
	return c.lockerRepo.GetAvailableLockers()
}
