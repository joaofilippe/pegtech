package services

import (
	"github.com/joaofilippe/pegtech/internal/domain/iservices"
	lockerusecases "github.com/joaofilippe/pegtech/internal/domain/usecases/locker"
)

// LockerReleaseService implements the LockerReleaseService interface
type LockerReleaseService struct {
	releaseLockerCase *lockerusecases.ReleaseLockerCase
}

// NewLockerReleaseService creates a new instance of LockerReleaseService
func NewLockerReleaseService(releaseLockerCase *lockerusecases.ReleaseLockerCase) iservices.LockerReleaseService {
	return &LockerReleaseService{
		releaseLockerCase: releaseLockerCase,
	}
}

// ReleaseLocker releases a locker by clearing its package information and setting it to available
func (s *LockerReleaseService) ReleaseLocker(id int) error {
	return s.releaseLockerCase.Execute(id)
}
