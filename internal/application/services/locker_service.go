package services

import (
	"github.com/google/uuid"
	"github.com/joaofilippe/pegtech/internal/domain/entities"
	irepositories "github.com/joaofilippe/pegtech/internal/domain/irepositories"
	"github.com/joaofilippe/pegtech/internal/domain/iservices"
	lockerusecases "github.com/joaofilippe/pegtech/internal/domain/usecases/locker"
)

type LockerService struct {
	registerLockerCase     *lockerusecases.RegisterLockerCase
	registerPackageCase    *lockerusecases.RegisterPackageCase
	reserveLockerCase      *lockerusecases.ReserveLockerCase
	getAvailableLockerCase *lockerusecases.GetAvailableLockersCase
	getLockerCase          *lockerusecases.GetLockerCase
	updateLockerStatusCase *lockerusecases.UpdateLockerStatusCase
	listLockersCase        *lockerusecases.ListLockersCase
	releaseLockerCase      *lockerusecases.ReleaseLockerCase
}

func NewLockerService(lockerRepo irepositories.LockerRepository) iservices.LockerService {
	return &LockerService{
		registerLockerCase:     lockerusecases.NewRegisterLockerCase(lockerRepo),
		registerPackageCase:    lockerusecases.NewRegisterPackageCase(lockerRepo),
		getAvailableLockerCase: lockerusecases.NewGetAvailableLockersCase(lockerRepo),
		getLockerCase:          lockerusecases.NewGetLockerCase(lockerRepo),
		updateLockerStatusCase: lockerusecases.NewUpdateLockerStatusCase(lockerRepo),
		listLockersCase:        lockerusecases.NewListLockersCase(lockerRepo),
		reserveLockerCase:      lockerusecases.NewReserveLockerCase(lockerRepo),
		releaseLockerCase:      lockerusecases.NewReleaseLockerCase(lockerRepo),
	}
}

// GetAvailableLockers implements iservices.LockerService.
func (s *LockerService) GetAvailableLockers() ([]int, error) {
	return s.getAvailableLockerCase.Execute()
}

func (s *LockerService) RegisterLocker(id int) error {
	return s.registerLockerCase.Execute(id)
}

func (s *LockerService) GetLocker(id int) (*entities.Locker, error) {
	return s.getLockerCase.Execute(id)
}

func (s *LockerService) UpdateLockerStatus(id int, status entities.LockerStatus) error {
	return s.updateLockerStatusCase.Execute(id, status)
}

func (s *LockerService) ListLockers() ([]*entities.Locker, error) {
	lockers, err := s.listLockersCase.Execute()
	return lockers, err
}

// ReserveLocker implements iservices.LockerService.
func (s *LockerService) ReserveLocker(userID uuid.UUID, expirationTime int) (string, error) {
	return s.reserveLockerCase.Execute(userID, expirationTime)
}

// RegisterPackage implements iservices.LockerService.
func (s *LockerService) RegisterPackage(userID uuid.UUID, expirationTime int) (string, error) {
	return s.registerPackageCase.Execute(userID, expirationTime)
}

func (s *LockerService) ReleaseLocker(id int) error {
	return s.releaseLockerCase.Execute(id)
}
