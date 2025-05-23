package services

import (
	"github.com/joaofilippe/pegtech/internal/domain/entities"
	irepositories "github.com/joaofilippe/pegtech/internal/domain/irepositories"
	"github.com/joaofilippe/pegtech/internal/domain/iservices"
	lockerusecases "github.com/joaofilippe/pegtech/internal/domain/usecases/locker"
)

type LockerService struct {
	registerLockerCase     *lockerusecases.RegisterLockerCase
	getAvailableLockerCase *lockerusecases.GetAvailableLockersCase
	getLockerCase          *lockerusecases.GetLockerCase
	updateLockerStatusCase *lockerusecases.UpdateLockerStatusCase
	listLockersCase        *lockerusecases.ListLockersCase
}

// OpenLocker implements iservices.LockerService.
func (s *LockerService) OpenLocker(lockerID int, password string) error {
	panic("unimplemented")
}

// GetAvailableLockers implements iservices.LockerService.
func (s *LockerService) GetAvailableLockers() ([]int, error) {
	panic("unimplemented")
}

func NewLockerService(lockerRepo irepositories.LockerRepository) iservices.LockerService {
	return &LockerService{
		registerLockerCase:     lockerusecases.NewRegisterLockerCase(lockerRepo),
		getAvailableLockerCase: lockerusecases.NewGetAvailableLockersCase(lockerRepo),
		getLockerCase:          lockerusecases.NewGetLockerCase(lockerRepo),
		updateLockerStatusCase: lockerusecases.NewUpdateLockerStatusCase(lockerRepo),
		listLockersCase:        lockerusecases.NewListLockersCase(lockerRepo),
	}
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
