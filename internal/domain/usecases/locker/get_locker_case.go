package lockerusecases

import (
	"github.com/joaofilippe/pegtech/internal/domain/entities"
	"github.com/joaofilippe/pegtech/internal/domain/irepositories"
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
func (uc *GetLockerCase) Execute(id int) (*entities.Port, error) {
	return uc.lockerRepo.GetLocker(id)
}
