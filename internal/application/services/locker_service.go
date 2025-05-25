package services

import (
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"github.com/joaofilippe/pegtech/internal/domain/entities"
	irepositories "github.com/joaofilippe/pegtech/internal/domain/irepositories"
	"github.com/joaofilippe/pegtech/internal/domain/iservices"
	lockerusecases "github.com/joaofilippe/pegtech/internal/domain/usecases/locker"
	"github.com/joaofilippe/pegtech/internal/infra/repositories/mqtt"
)

type LockerService struct {
	registerLockerCase     *lockerusecases.RegisterLockerCase
	registerPackageCase    *lockerusecases.RegisterPackageCase
	reserveLockerCase      *lockerusecases.ReserveLockerCase
	getAvailableLockerCase *lockerusecases.GetAvailableLockersCase
	getLockerCase          *lockerusecases.GetLockerCase
	updateLockerStatusCase *lockerusecases.UpdatePortStatusCase
	listLockersCase        *lockerusecases.ListLockersCase
	releaseLockerCase      *lockerusecases.ReleaseLockerCase
	pickupPackageCase      *lockerusecases.PickupPackageCase
	getPackagesByUserCase  *lockerusecases.GetPackagesByUserCase
	mqttClient             *mqtt.MqttClient
}

func NewLockerService(lockerRepo irepositories.LockerRepository) iservices.LockerService {
	return &LockerService{
		registerLockerCase:     lockerusecases.NewRegisterLockerCase(lockerRepo),
		registerPackageCase:    lockerusecases.NewRegisterPackageCase(lockerRepo),
		getAvailableLockerCase: lockerusecases.NewGetAvailableLockersCase(lockerRepo),
		getLockerCase:          lockerusecases.NewGetLockerCase(lockerRepo),
		updateLockerStatusCase: lockerusecases.NewUpdatePortStatusCase(lockerRepo),
		listLockersCase:        lockerusecases.NewListLockersCase(lockerRepo),
		reserveLockerCase:      lockerusecases.NewReserveLockerCase(lockerRepo),
		releaseLockerCase:      lockerusecases.NewReleaseLockerCase(lockerRepo),
		pickupPackageCase:      lockerusecases.NewPickupPackageCase(lockerRepo),
		getPackagesByUserCase:  lockerusecases.NewGetPackagesByUserCase(lockerRepo),
		mqttClient:             lockerRepo.GetMQTTClient(),
	}
}

// GetAvailableLockers implements iservices.LockerService.
func (s *LockerService) GetAvailableLockers() ([]int, error) {
	return s.getAvailableLockerCase.Execute()
}

func (s *LockerService) RegisterLocker(lockerID int, ports []int) error {
	return s.registerLockerCase.Execute(lockerID, ports)
}

func (s *LockerService) GetLocker(id int) (*entities.Port, error) {
	return s.getLockerCase.Execute(id)
}

func (s *LockerService) UpdatePortStatus(id string, status entities.LockerStatus) error {
	return s.updateLockerStatusCase.Execute(id, status)
}

func (s *LockerService) ListLockers() ([]*entities.Port, error) {
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

// ReleaseLocker implements iservices.LockerService
func (s *LockerService) ReleaseLocker(lockerID int, packageCode string) error {
	return s.releaseLockerCase.Execute(lockerID, packageCode)
}

// PickupPackage implements iservices.LockerService
func (s *LockerService) PickupPackage(packageCode string, password string) error {
	return s.pickupPackageCase.Execute(packageCode, password)
}

// StartPackagePickupSubscription starts listening to package pickup events
func (s *LockerService) StartPackagePickupSubscription() (chan []byte, error) {
	subscriber := mqtt.NewSubscriber(s.mqttClient)

	lockerChan, err := subscriber.SubscribeToPackagePickup()
	if err != nil {
		return nil, err
	}

	// Inicia uma goroutine para processar os IDs dos lockers
	go func() {
		for locker := range lockerChan {
			var lockerInput struct {
				ID          int    `json:"locker_id"`
				PackageCode string `json:"package_code"`
			}

			if err := json.Unmarshal(locker, &lockerInput); err != nil {
				log.Printf("Erro ao decodificar mensagem MQTT: %v", err)
				continue
			}

			// Libera o locker
			if err := s.ReleaseLocker(lockerInput.ID, lockerInput.PackageCode); err != nil {
				log.Printf("Erro ao liberar locker %d: %v", lockerInput.ID, err)
				continue
			}
			log.Printf("Locker %d liberado após retirada do pacote", lockerInput.ID)
		}
	}()

	return lockerChan, nil
}

// StartRegisterPackageSubscription implements iservices.LockerService.
func (s *LockerService) StartRegisterPackageSubscription() error {
	subscriber := mqtt.NewSubscriber(s.mqttClient)

	registerChan, err := subscriber.SubscribeToPackageRegistration()
	if err != nil {
		return err
	}

	go func() {
		for register := range registerChan {
			var registerInput struct {
				ID          int    `json:"locker_id"`
				PackageCode string `json:"package_code"`
				Message     string `json:"msg"`
			}

			if err := json.Unmarshal(register, &registerInput); err != nil {
				log.Printf("Erro ao decodificar mensagem MQTT: %v", err)
				continue
			}

			if len(registerInput.Message) > 0 {
				log.Printf("Pacote %s registrado no locker %d", registerInput.PackageCode, registerInput.ID)
				s.UpdatePortStatus(registerInput.PackageCode, entities.LockerStatusOccupied)
			}

		}
	}()

	return nil
}

func (s *LockerService) StartLockerAvailableSubscription() (chan []byte, error) {
	subscriber := mqtt.NewSubscriber(s.mqttClient)

	availableChan, err := subscriber.SubscribeToLockerAvailable()
	if err != nil {
		return nil, err
	}

	for available := range availableChan {
		var availableData struct {
			Lockers []entities.Port `json:"lockers"`
		}

		if err := json.Unmarshal(available, &availableData); err != nil {
			log.Printf("Erro ao decodificar mensagem MQTT: %v", err)
			continue
		}
	}

	return availableChan, nil
}

// GetPackagesByUser implements iservices.LockerService
func (s *LockerService) GetPackagesByUser(userID uuid.UUID) ([]*entities.Port, error) {
	return s.getPackagesByUserCase.Execute(userID)
}
