package lockerusecases

import (
	"github.com/joaofilippe/pegtech/domain/entities"
	"github.com/joaofilippe/pegtech/domain/irepositories"
)

// GetLockerCase handles locker retrieval by ID
type GetLockerCase struct {
	lockerRepo irepositories.LockerRepository
}

// NewGetLockerCase creates a new instance of GetLockerCase
func NewGetLockerCase(lockerRepo irepositories.LockerRepository) *GetLockerCase {
	return &GetLockerCase{
		lockerRepo: lockerRepo,
	}
}

// Execute performs the locker retrieval operation
func (uc *GetLockerCase) Execute(id string) (*entities.Locker, error) {
	return uc.lockerRepo.GetLocker(id)
}
