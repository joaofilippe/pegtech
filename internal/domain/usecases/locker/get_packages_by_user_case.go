package lockerusecases

import (
	"github.com/google/uuid"
	"github.com/joaofilippe/pegtech/internal/domain/entities"
	"github.com/joaofilippe/pegtech/internal/domain/irepositories"
)

// GetPackagesByUserCase handles retrieving packages for a specific user
type GetPackagesByUserCase struct {
	lockerRepo irepositories.LockerRepository
}

// NewGetPackagesByUserCase creates a new instance of GetPackagesByUserCase
func NewGetPackagesByUserCase(lockerRepo irepositories.LockerRepository) *GetPackagesByUserCase {
	return &GetPackagesByUserCase{
		lockerRepo: lockerRepo,
	}
}

// Execute performs the package retrieval operation
func (uc *GetPackagesByUserCase) Execute(userID uuid.UUID) ([]*entities.Port, error) {
	return uc.lockerRepo.GetPackagesByUser(userID)
}
