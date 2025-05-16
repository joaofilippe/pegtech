package memory

import (
	"errors"
	"sync"

	"github.com/joaofilippe/pegtech/domain/entities"
)

type PackageRepository struct {
	packages map[string]*entities.Package
	mu       sync.RWMutex
}

func NewPackageRepository() *PackageRepository {
	return &PackageRepository{
		packages: make(map[string]*entities.Package),
	}
}

func (r *PackageRepository) SavePackage(pkg *entities.Package) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.packages[pkg.ID] = pkg
	return nil
}

func (r *PackageRepository) GetPackageByTrackingCode(trackingCode string) (*entities.Package, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, pkg := range r.packages {
		if pkg.TrackingCode == trackingCode {
			return pkg, nil
		}
	}
	return nil, errors.New("package not found")
}

func (r *PackageRepository) DeletePackage(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.packages[id]; !exists {
		return errors.New("package not found")
	}

	delete(r.packages, id)
	return nil
}
