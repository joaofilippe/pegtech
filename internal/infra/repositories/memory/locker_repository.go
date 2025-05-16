package memory

import (
	"errors"
	"sync"

	"github.com/joaofilippe/pegtech/domain/entities"
)

type LockerRepository struct {
	lockers map[string]*entities.Locker
	mu      sync.RWMutex
}

func NewLockerRepository() *LockerRepository {
	return &LockerRepository{
		lockers: make(map[string]*entities.Locker),
	}
}

func (r *LockerRepository) SaveLocker(locker *entities.Locker) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.lockers[locker.ID] = locker
	return nil
}

func (r *LockerRepository) GetLocker(id string) (*entities.Locker, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	locker, exists := r.lockers[id]
	if !exists {
		return nil, errors.New("locker not found")
	}
	return locker, nil
}

func (r *LockerRepository) GetAvailableLocker(size string) (*entities.Locker, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, locker := range r.lockers {
		if locker.Status == entities.LockerAvailable && locker.Size == size {
			return locker, nil
		}
	}
	return nil, errors.New("no available lockers")
}

func (r *LockerRepository) UpdateLockerStatus(id string, status entities.LockerStatus) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	locker, exists := r.lockers[id]
	if !exists {
		return errors.New("locker not found")
	}

	locker.Status = status
	return nil
}

func (r *LockerRepository) ListLockers() ([]*entities.Locker, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	lockers := make([]*entities.Locker, 0, len(r.lockers))
	for _, locker := range r.lockers {
		lockers = append(lockers, locker)
	}
	return lockers, nil
}
