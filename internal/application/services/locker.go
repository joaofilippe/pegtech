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
	getAvailableLockerCase *lockerusecases.GetAvailableLockerCase
	getLockerCase          *lockerusecases.GetLockerCase
	updateLockerStatusCase *lockerusecases.UpdateLockerStatusCase
	registerPackageCase    *lockerusecases.RegisterPackageCase
	openLockerCase         *lockerusecases.OpenLockerCase
	listLockersCase        *lockerusecases.ListLockersCase
}

func NewLockerService(lockerRepo irepositories.LockerRepository, packageRepo irepositories.PackageRepository) iservices.LockerService {
	return &LockerService{
		registerLockerCase:     lockerusecases.NewRegisterLockerCase(lockerRepo),
		getAvailableLockerCase: lockerusecases.NewGetAvailableLockerCase(lockerRepo),
		getLockerCase:          lockerusecases.NewGetLockerCase(lockerRepo),
		updateLockerStatusCase: lockerusecases.NewUpdateLockerStatusCase(lockerRepo),
		registerPackageCase:    lockerusecases.NewRegisterPackageCase(lockerRepo, packageRepo),
		openLockerCase:         lockerusecases.NewOpenLockerCase(lockerRepo, packageRepo),
		listLockersCase:        lockerusecases.NewListLockersCase(lockerRepo),
	}
}

func (s *LockerService) RegisterLocker(id int) error {
	return s.registerLockerCase.Execute(id)
}

func (s *LockerService) GetAvailableLocker(size string) (*entities.Locker, error) {
	return s.getAvailableLockerCase.Execute(size)
}

func (s *LockerService) GetLocker(id int) (*entities.Locker, error) {
	return s.getLockerCase.Execute(id)
}

func (s *LockerService) UpdateLockerStatus(id int, status entities.LockerStatus) error {
	return s.updateLockerStatusCase.Execute(id, status)
}

func (s *LockerService) RegisterPackage(userID uuid.UUID, lockerID, expirationTime int) (*entities.Package, error) {
	return s.registerPackageCase.Execute(userID, lockerID, expirationTime)
}

func (s *LockerService) OpenLocker(lockerID int, password string) error {
	return s.openLockerCase.Execute(lockerID, password)
}

func (s *LockerService) ListLockers() ([]*entities.Locker, error) {
	lockers, err := s.listLockersCase.Execute()
	return lockers, err
}
