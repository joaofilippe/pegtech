package services

import (
	"github.com/joaofilippe/pegtech/domain/entities"
	irepositories "github.com/joaofilippe/pegtech/domain/irepositories"
	"github.com/joaofilippe/pegtech/domain/iservices"
	lockerusecases "github.com/joaofilippe/pegtech/domain/usecases/locker"
)

type LockerService struct {
	registerLockerCase     *lockerusecases.RegisterLockerCase
	getAvailableLockerCase *lockerusecases.GetAvailableLockerCase
	registerPackageCase    *lockerusecases.RegisterPackageCase
	getPackagePickupCase   *lockerusecases.GetPackagePickupInfoCase
	openLockerCase         *lockerusecases.OpenLockerCase
	listLockersCase        *lockerusecases.ListLockersCase
}

func NewLockerService(lockerRepo irepositories.LockerRepository, packageRepo irepositories.PackageRepository) iservices.LockerService {
	return &LockerService{
		registerLockerCase:     lockerusecases.NewRegisterLockerCase(lockerRepo),
		getAvailableLockerCase: lockerusecases.NewGetAvailableLockerCase(lockerRepo),
		registerPackageCase:    lockerusecases.NewRegisterPackageCase(lockerRepo, packageRepo),
		getPackagePickupCase:   lockerusecases.NewGetPackagePickupInfoCase(packageRepo),
		openLockerCase:         lockerusecases.NewOpenLockerCase(lockerRepo, packageRepo),
		listLockersCase:        lockerusecases.NewListLockersCase(lockerRepo),
	}
}

func (s *LockerService) RegisterLocker(id string, size string) error {
	return s.registerLockerCase.Execute(id, size)
}

func (s *LockerService) GetAvailableLocker(size string) (*entities.Locker, error) {
	return s.getAvailableLockerCase.Execute(size)
}

func (s *LockerService) RegisterPackage(trackingCode string, size string) (*entities.Package, error) {
	return s.registerPackageCase.Execute(trackingCode, size)
}

func (s *LockerService) GetPackagePickupInfo(trackingCode string) (*entities.PackagePickup, error) {
	return s.getPackagePickupCase.Execute(trackingCode)
}

func (s *LockerService) OpenLocker(lockerID string, password string) error {
	return s.openLockerCase.Execute(lockerID, password)
}

func (s *LockerService) ListLockers() ([]*entities.Locker, error) {
	return s.listLockersCase.Execute()
}
