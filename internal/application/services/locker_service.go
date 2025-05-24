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
	updateLockerStatusCase *lockerusecases.UpdateLockerStatusCase
	listLockersCase        *lockerusecases.ListLockersCase
	releaseLockerCase      *lockerusecases.ReleaseLockerCase
	mqttClient             *mqtt.MqttClient
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
		mqttClient:             lockerRepo.GetMQTTClient(),
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

// ReleaseLocker implements iservices.LockerService
func (s *LockerService) ReleaseLocker(lockerID int, packageCode string) error {
	return s.releaseLockerCase.Execute(lockerID, packageCode)
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
				ID int `json:"locker_id"`
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

	err := subscriber.SubscribeToPackageRegistration()
	if err != nil {
		return err
	}

	_, err = subscriber.SubscribeToLockerAvailable()
	if err != nil {
		return err
	}
	

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
			Lockers []entities.Locker `json:"lockers"`
		}

		if err := json.Unmarshal(available, &availableData); err != nil {
			log.Printf("Erro ao decodificar mensagem MQTT: %v", err)
			continue
		}
	}

	return availableChan, nil
}
