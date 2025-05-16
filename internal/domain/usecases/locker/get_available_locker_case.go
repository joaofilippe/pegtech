package lockerusecases

import (
	"github.com/joaofilippe/pegtech/domain/entities"
	"github.com/joaofilippe/pegtech/domain/irepositories"
)

// GetAvailableLockerCase handles retrieving available lockers
type GetAvailableLockerCase struct {
	lockerRepo irepositories.LockerRepository
}

// NewGetAvailableLockerCase creates a new instance of GetAvailableLockerCase
func NewGetAvailableLockerCase(lockerRepo irepositories.LockerRepository) *GetAvailableLockerCase {
	return &GetAvailableLockerCase{
		lockerRepo: lockerRepo,
	}
}

// Execute performs the available locker retrieval operation
func (uc *GetAvailableLockerCase) Execute(size string) (*entities.Locker, error) {
	return uc.lockerRepo.GetAvailableLocker(size)
}
