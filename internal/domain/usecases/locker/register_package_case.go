package lockerusecases

import (
	"sort"
	"time"

	"github.com/google/uuid"
	"github.com/joaofilippe/pegtech/internal/domain/entities"
	"github.com/joaofilippe/pegtech/internal/domain/errors"
	"github.com/joaofilippe/pegtech/internal/domain/irepositories"
)

// RegisterPackageCase handles package registration
type RegisterPackageCase struct {
	lockerRepo irepositories.LockerRepository
}

// NewRegisterPackageCase creates a new instance of RegisterPackageCase
func NewRegisterPackageCase(lockerRepo irepositories.LockerRepository) *RegisterPackageCase {
	return &RegisterPackageCase{
		lockerRepo: lockerRepo,
	}
}

// Execute performs the package registration operation
func (uc *RegisterPackageCase) Execute(userID uuid.UUID, expirationTime int) (string, error) {
	ports, err := uc.lockerRepo.ListLockers()
	if err != nil {
		return "", err
	}

	if len(ports) == 0 {
		return "", errors.ErrFoundNoLockers
	}

	// Find an available locker
	availableLocker, err := uc.getAvailablePort(ports)
	if err != nil {
		return "", err
	}

	if availableLocker == nil {
		return "", errors.ErrNoAvailableLockers
	}

	// Generate package code and password
	packageCode, err := generatePassword()
	if err != nil {
		return "", err
	}
	packagePassword, err := generatePassword()
	if err != nil {
		return "", err
	}

	// Calculate expiration time
	expiresAt := time.Now().Add(time.Duration(expirationTime) * time.Hour)

	// Create package registration
	registration := irepositories.PackageRegistration{
		PackageCode:           packageCode,
		PackagePickupPassword: packagePassword,
		UserID:                userID,
		ExpiresAt:             &expiresAt,
	}

	// Register the package in the locker
	err = uc.lockerRepo.RegisterPackage(availableLocker.ID, registration)
	if err != nil {
		return "", err
	}

	// Update locker status to occupied
	err = uc.lockerRepo.UpdateLockerStatus(availableLocker.ID, entities.LockerStatusOccupied)
	if err != nil {
		return "", err
	}

	return packageCode, nil
}

func (uc *RegisterPackageCase) getAvailablePort(ports []*entities.Port) (*entities.Port, error) {
	availablePorts := make([]*entities.Port, 0)
	for _, port := range ports {
		if port.Status == entities.LockerStatusAvailable {
			availablePorts = append(availablePorts, port)
		}
	}

	if len(availablePorts) == 0 {
		return nil, ErrNoAvailablePorts
	}

	// Sort ports first by locker number, then by port number
	sort.Slice(availablePorts, func(i, j int) bool {
		// If lockers are different, sort by locker
		if availablePorts[i].Locker != availablePorts[j].Locker {
			return availablePorts[i].Locker < availablePorts[j].Locker
		}
		// If same locker, sort by port number
		return availablePorts[i].Port < availablePorts[j].Port
	})

	availablePort := availablePorts[0]
	contains := false
	for _, port := range availablePorts {
		locker, err := uc.lockerRepo.GetAvailablePorts(availablePort.Locker)
		if err != nil {
			return nil, err
		}

		for _, lockerPort := range locker.Ports {
			if lockerPort.Port == port.Port {
				contains = true
				break
			}
		}

		if contains {
			availablePort = port
			break
		}
	}

	if !contains {
		return nil, ErrNoAvailablePorts
	}

	return availablePort, nil
}
