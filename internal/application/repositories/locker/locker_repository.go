package repositories

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/joaofilippe/pegtech/internal/domain/entities"
	"github.com/joaofilippe/pegtech/internal/domain/irepositories"
	"github.com/joaofilippe/pegtech/internal/infra/repositories/database"
	"github.com/joaofilippe/pegtech/internal/infra/repositories/mqtt"
)

var (
	ErrLockerNotFound = errors.New("locker not found")
	ErrPortNotFound   = errors.New("port not found")
)

// LockerRepository implements the LockerRepository interface
type LockerRepository struct {
	db         *database.PostgresDB
	mqttClient *mqtt.MqttClient
	subscriber *mqtt.Subscriber
}

// NewLockerRepository creates a new instance of LockerRepository
func NewLockerRepository(db *database.PostgresDB, mqttClient *mqtt.MqttClient) irepositories.LockerRepository {
	return &LockerRepository{
		db:         db,
		mqttClient: mqttClient,
		subscriber: mqtt.NewSubscriber(mqttClient),
	}
}

// SavePort saves a port to the storage
func (r *LockerRepository) SavePort(port *entities.Port) error {
	_, err := r.db.DB().Exec(SavePortQuery,
		port.ID,
		port.Locker,
		port.Port,
		port.Number,
		port.PackageCode,
		port.PackagePickupPassword,
		port.PackagePickupExpiresAt,
		port.PackageUserID,
		port.Status,
		port.ReservedExpiration,
		port.OccupiedAt,
		port.OccupiedUntil,
		port.CreatedAt,
		port.UpdatedAt,
	)
	return err
}

// GetPort retrieves a port by ID
func (r *LockerRepository) GetPort(lockerID, portID int) (*entities.Port, error) {
	port := &entities.Port{}
	err := r.db.DB().QueryRow(GetPortQuery, portID, lockerID).Scan(
		&port.ID,
		&port.Locker,
		&port.Port,
		&port.Number,
		&port.PackageCode,
		&port.PackagePickupPassword,
		&port.PackagePickupExpiresAt,
		&port.PackageUserID,
		&port.Status,
		&port.ReservedExpiration,
		&port.OccupiedAt,
		&port.OccupiedUntil,
		&port.CreatedAt,
		&port.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, ErrPortNotFound
	}

	if err != nil {
		return nil, err
	}

	return port, nil
}

// UpdatePortStatus updates the status of a port
func (r *LockerRepository) UpdatePortStatus(lockerID, portID int, status entities.LockerStatus) error {
	result, err := r.db.DB().Exec(UpdatePortStatusQuery, status, time.Now(), portID, lockerID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrPortNotFound
	}

	return nil
}

// ListPorts retrieves all ports for a locker
func (r *LockerRepository) ListPorts(lockerID int) ([]*entities.Port, error) {
	rows, err := r.db.DB().Query(ListPortsQuery, lockerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ports []*entities.Port
	for rows.Next() {
		port := &entities.Port{}
		err := rows.Scan(
			&port.ID,
			&port.Locker,
			&port.Port,
			&port.Number,
			&port.PackageCode,
			&port.PackagePickupPassword,
			&port.PackagePickupExpiresAt,
			&port.PackageUserID,
			&port.Status,
			&port.ReservedExpiration,
			&port.OccupiedAt,
			&port.OccupiedUntil,
			&port.CreatedAt,
			&port.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		ports = append(ports, port)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return ports, nil
}

// GetAvailablePorts retrieves all available ports for a locker
func (r *LockerRepository) GetAvailablePorts(lockerID int) (entities.Locker, error) {
	// Criar canal temporário para esta consulta específica
	tempChan := make(chan []byte, 10)

	// Função para fechar o canal de forma segura
	defer func() {
		select {
		case <-tempChan:
		default:
		}
		close(tempChan)
	}()

	// Subscribe temporariamente usando o cliente MQTT diretamente
	err := r.mqttClient.Subscribe("locker/available/joao", func(payload []byte) {
		select {
		case tempChan <- payload:
		default:
			// Canal cheio, ignora mensagem
		}
	})

	if err != nil {
		return entities.Locker{}, err
	}

	// Publicar requisição
	requestMap := map[string]int{"locker_id": lockerID}
	if err := r.mqttClient.Publish("locker/available/joao", requestMap); err != nil {
		log.Printf("Error requesting available ports from MQTT: %v", err)
		return entities.Locker{}, err
	}

	locker := entities.Locker{
		ID:    lockerID,
		Ports: make([]*entities.Port, 0),
	}

	// Timeout
	timeout := time.After(15 * time.Second)

	for {
		select {
		case available := <-tempChan:
			var availableData struct {
				LockerID int   `json:"locker_id"`
				Ports    []int `json:"ports"`
			}

			if err := json.Unmarshal(available, &availableData); err != nil {
				log.Printf("Error unmarshalling available ports: %v", err)
				continue
			}

			if availableData.LockerID == lockerID && len(availableData.Ports) > 0 {
				for _, portNum := range availableData.Ports {
					port := &entities.Port{
						Locker: lockerID,
						Port:   portNum,
						Status: entities.LockerStatusAvailable,
					}
					locker.Ports = append(locker.Ports, port)
				}
				return locker, nil
			}

		case <-timeout:
			log.Printf("Timeout waiting for MQTT response for locker %d", lockerID)
			return locker, nil
		}
	}
}

// RegisterPackage registers a package in a port
func (r *LockerRepository) RegisterPackage(lockerID int, registration irepositories.PackageRegistration) error {
	packageMQTT := struct {
		LockerID              int        `json:"locker_id"`
		Port                  int        `json:"port"`
		PackageCode           string     `json:"package_code"`
		PackagePickupPassword string     `json:"package_pickup_password"`
		ExpiresAt             *time.Time `json:"expires_at"`
	}{
		LockerID:              lockerID,
		Port:                  lockerID, // Using lockerID as port for now
		PackageCode:           registration.PackageCode,
		PackagePickupPassword: registration.PackagePickupPassword,
		ExpiresAt:             registration.ExpiresAt,
	}

	err := r.mqttClient.Publish("locker/package/register/joao", packageMQTT)
	if err != nil {
		return err
	}

	_, err = r.db.DB().Exec(RegisterPackageQuery,
		registration.PackageCode,
		registration.PackagePickupPassword,
		registration.UserID,
		registration.ExpiresAt,
		time.Now(),
		"RESERVED",
		lockerID,
	)
	return err
}

// ReservePort reserves a port for a user
func (r *LockerRepository) ReservePort(lockerID, portID int, userID uuid.UUID, expiration *time.Time) error {
	_, err := r.db.DB().Exec(ReservePortQuery,
		userID,
		expiration,
		time.Now(),
		portID,
		lockerID,
	)
	return err
}

// UpdatePort updates a port in the storage
func (r *LockerRepository) UpdatePort(port *entities.Port) error {
	_, err := r.db.DB().Exec(UpdatePortQuery,
		port.Number,
		port.PackageCode,
		port.PackagePickupPassword,
		port.PackagePickupExpiresAt,
		port.PackageUserID,
		port.Status,
		port.ReservedExpiration,
		port.OccupiedAt,
		port.OccupiedUntil,
		port.UpdatedAt,
		port.ID,
		port.Locker,
	)
	return err
}

// ReleasePort releases a port by clearing its package information and setting it to available
func (r *LockerRepository) ReleasePort(lockerID, portID int) error {
	_, err := r.db.DB().Exec(ReleasePortQuery,
		"",                             // package_code
		"",                             // package_pickup_password
		nil,                            // package_pickup_expires_at
		nil,                            // package_user_id
		entities.LockerStatusAvailable, // status
		nil,                            // reserved_expiration
		nil,                            // occupied_at
		nil,                            // occupied_until
		time.Now(),                     // updated_at
		portID,
		lockerID,
	)
	return err
}

// GetLocker retrieves a locker by ID
func (r *LockerRepository) GetLocker(id int) (*entities.Port, error) {
	port := &entities.Port{}
	err := r.db.DB().QueryRow(GetPortQuery, id, 0).Scan(
		&port.ID,
		&port.Locker,
		&port.Port,
		&port.Number,
		&port.PackageCode,
		&port.PackagePickupPassword,
		&port.PackagePickupExpiresAt,
		&port.PackageUserID,
		&port.Status,
		&port.ReservedExpiration,
		&port.OccupiedAt,
		&port.OccupiedUntil,
		&port.CreatedAt,
		&port.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, ErrLockerNotFound
	}

	if err != nil {
		return nil, err
	}

	return port, nil
}

// ListLockers retrieves all lockers
func (r *LockerRepository) ListLockers() ([]*entities.Port, error) {
	rows, err := r.db.DB().Query(ListPortsQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ports []*entities.Port
	for rows.Next() {
		port := &entities.Port{}
		err := rows.Scan(
			&port.ID,
			&port.Locker,
			&port.Port,
			&port.Number,
			&port.PackageCode,
			&port.PackagePickupPassword,
			&port.PackagePickupExpiresAt,
			&port.PackageUserID,
			&port.Status,
			&port.ReservedExpiration,
			&port.OccupiedAt,
			&port.OccupiedUntil,
			&port.CreatedAt,
			&port.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		ports = append(ports, port)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return ports, nil
}

// ReleaseLocker releases a locker by clearing its package information and setting it to available
func (r *LockerRepository) ReleaseLocker(lockerID int, packageCode string) error {
	_, err := r.db.DB().Exec(ReleasePortQuery,
		"",                             // package_code
		"",                             // package_pickup_password
		nil,                            // package_pickup_expires_at
		nil,                            // package_user_id
		entities.LockerStatusAvailable, // status
		nil,                            // reserved_expiration
		nil,                            // occupied_at
		nil,                            // occupied_until
		time.Now(),                     // updated_at
		lockerID,                       // port_id
		lockerID,                       // locker_id
	)
	return err
}

// SaveLocker saves a locker to the storage
func (r *LockerRepository) SaveLocker(locker *entities.Port) error {
	_, err := r.db.DB().Exec(SavePortQuery,
		locker.Locker,
		locker.Port,
		locker.Number,
		locker.PackageCode,
		locker.PackagePickupPassword,
		locker.PackagePickupExpiresAt,
		locker.PackageUserID,
		locker.Status,
		locker.ReservedExpiration,
		locker.OccupiedAt,
		locker.OccupiedUntil,
		locker.CreatedAt,
		locker.UpdatedAt,
	)
	return err
}

// UpdateLocker updates a locker in the storage
func (r *LockerRepository) UpdateLocker(locker *entities.Port) error {
	_, err := r.db.DB().Exec(UpdatePortQuery,
		locker.Number,
		locker.PackageCode,
		locker.PackagePickupPassword,
		locker.PackagePickupExpiresAt,
		locker.PackageUserID,
		locker.Status,
		locker.ReservedExpiration,
		locker.OccupiedAt,
		locker.OccupiedUntil,
		locker.UpdatedAt,
		locker.ID,
		locker.Locker,
	)
	return err
}

// UpdateLockerStatus updates the status of a locker
func (r *LockerRepository) UpdateLockerStatus(id int, status entities.LockerStatus) error {
	result, err := r.db.DB().Exec(UpdatePortStatusQuery, status, time.Now(), id, 0)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrLockerNotFound
	}

	return nil
}

// GetMQTTClient returns the MQTT client
func (r *LockerRepository) GetMQTTClient() *mqtt.MqttClient {
	return r.mqttClient
}

// GetPackagesByUser retrieves all packages for a specific user
func (r *LockerRepository) GetPackagesByUser(userID uuid.UUID) ([]*entities.Port, error) {
	rows, err := r.db.DB().Query(GetPackagesByUserQuery, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ports []*entities.Port
	for rows.Next() {
		port := &entities.Port{}
		err := rows.Scan(
			&port.ID,
			&port.Locker,
			&port.Port,
			&port.Number,
			&port.PackageCode,
			&port.PackagePickupPassword,
			&port.PackagePickupExpiresAt,
			&port.PackageUserID,
			&port.Status,
			&port.ReservedExpiration,
			&port.OccupiedAt,
			&port.OccupiedUntil,
			&port.CreatedAt,
			&port.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		ports = append(ports, port)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return ports, nil
}

// GetLockerByPackageCode retrieves a locker by package code
func (r *LockerRepository) GetLockerByPackageCode(packageCode string) (*entities.Port, error) {
	port := &entities.Port{}
	err := r.db.DB().QueryRow(GetPortByPackageCodeQuery, packageCode).Scan(
		&port.ID,
		&port.Locker,
		&port.Port,
		&port.Number,
		&port.PackageCode,
		&port.PackagePickupPassword,
		&port.PackagePickupExpiresAt,
		&port.PackageUserID,
		&port.Status,
		&port.ReservedExpiration,
		&port.OccupiedAt,
		&port.OccupiedUntil,
		&port.CreatedAt,
		&port.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, ErrLockerNotFound
	}

	if err != nil {
		return nil, err
	}

	return port, nil
}
