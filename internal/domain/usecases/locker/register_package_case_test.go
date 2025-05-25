package lockerusecases

import (
	"testing"

	"github.com/google/uuid"
	"github.com/joaofilippe/pegtech/internal/domain/entities"
	"github.com/joaofilippe/pegtech/internal/domain/irepositories"
	"github.com/joaofilippe/pegtech/internal/infra/repositories/mqtt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockLockerRepository is a mock implementation of LockerRepository
type MockLockerRepository struct {
	mock.Mock
}

func (m *MockLockerRepository) SaveLocker(locker *entities.Port) error {
	args := m.Called(locker)
	return args.Error(0)
}

func (m *MockLockerRepository) GetLocker(id int) (*entities.Port, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.Port), args.Error(1)
}

func (m *MockLockerRepository) GetAvailablePorts(lockerID int) (entities.Locker, error) {
	args := m.Called(lockerID)
	return args.Get(0).(entities.Locker), args.Error(1)
}

func (m *MockLockerRepository) ListLockers() ([]*entities.Port, error) {
	args := m.Called()
	return args.Get(0).([]*entities.Port), args.Error(1)
}

func (m *MockLockerRepository) UpdateLocker(locker *entities.Port) error {
	args := m.Called(locker)
	return args.Error(0)
}

func (m *MockLockerRepository) RegisterPackage(lockerID int, registration irepositories.PackageRegistration) error {
	args := m.Called(lockerID, registration)
	return args.Error(0)
}

func (m *MockLockerRepository) UpdateLockerStatus(id int, status entities.LockerStatus) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func (m *MockLockerRepository) ReleaseLocker(lockerID int, packageCode string) error {
	args := m.Called(lockerID, packageCode)
	return args.Error(0)
}

func (m *MockLockerRepository) GetPackagesByUser(userID uuid.UUID) ([]*entities.Port, error) {
	args := m.Called(userID)
	return args.Get(0).([]*entities.Port), args.Error(1)
}

func (m *MockLockerRepository) GetMQTTClient() *mqtt.MqttClient {
	args := m.Called()
	return args.Get(0).(*mqtt.MqttClient)
}

func TestGetAvailablePorts(t *testing.T) {
	// Create test cases
	tests := []struct {
		name          string
		ports         []*entities.Port
		mockLocker    entities.Locker
		expectedPort  *entities.Port
		expectedError error
	}{
		{
			name: "should return first available port when multiple ports are available",
			ports: []*entities.Port{
				{
					ID:     1,
					Locker: 1,
					Port:   1,
					Status: entities.LockerStatusAvailable,
				},
				{
					ID:     2,
					Locker: 1,
					Port:   2,
					Status: entities.LockerStatusAvailable,
				},
			},
			mockLocker: entities.Locker{
				ID: 1,
				Ports: []*entities.Port{
					{
						ID:     1,
						Locker: 1,
						Port:   1,
						Status: entities.LockerStatusAvailable,
					},
				},
			},
			expectedPort: &entities.Port{
				ID:     1,
				Locker: 1,
				Port:   1,
				Status: entities.LockerStatusAvailable,
			},
			expectedError: nil,
		},
		{
			name: "should return error when no ports are available",
			ports: []*entities.Port{
				{
					ID:     1,
					Locker: 1,
					Port:   1,
					Status: entities.LockerStatusOccupied,
				},
			},
			expectedPort:  nil,
			expectedError: ErrNoAvailablePorts,
		},
		{
			name: "should sort by locker number and return port from lowest locker",
			ports: []*entities.Port{
				{
					ID:     3,
					Locker: 3,
					Port:   1,
					Status: entities.LockerStatusAvailable,
				},
				{
					ID:     1,
					Locker: 1,
					Port:   1,
					Status: entities.LockerStatusAvailable,
				},
				{
					ID:     2,
					Locker: 2,
					Port:   1,
					Status: entities.LockerStatusAvailable,
				},
			},
			mockLocker: entities.Locker{
				ID: 1,
				Ports: []*entities.Port{
					{
						ID:     1,
						Locker: 1,
						Port:   1,
						Status: entities.LockerStatusAvailable,
					},
				},
			},
			expectedPort: &entities.Port{
				ID:     1,
				Locker: 1,
				Port:   1,
				Status: entities.LockerStatusAvailable,
			},
			expectedError: nil,
		},
		{
			name: "should sort by port number within same locker",
			ports: []*entities.Port{
				{
					ID:     3,
					Locker: 1,
					Port:   3,
					Status: entities.LockerStatusAvailable,
				},
				{
					ID:     1,
					Locker: 1,
					Port:   1,
					Status: entities.LockerStatusAvailable,
				},
				{
					ID:     2,
					Locker: 1,
					Port:   2,
					Status: entities.LockerStatusAvailable,
				},
			},
			mockLocker: entities.Locker{
				ID: 1,
				Ports: []*entities.Port{
					{
						ID:     1,
						Locker: 1,
						Port:   1,
						Status: entities.LockerStatusAvailable,
					},
				},
			},
			expectedPort: &entities.Port{
				ID:     1,
				Locker: 1,
				Port:   1,
				Status: entities.LockerStatusAvailable,
			},
			expectedError: nil,
		},
		{
			name: "should handle mixed lockers and ports correctly",
			ports: []*entities.Port{
				{
					ID:     5,
					Locker: 2,
					Port:   2,
					Status: entities.LockerStatusAvailable,
				},
				{
					ID:     3,
					Locker: 1,
					Port:   3,
					Status: entities.LockerStatusAvailable,
				},
				{
					ID:     1,
					Locker: 1,
					Port:   1,
					Status: entities.LockerStatusAvailable,
				},
				{
					ID:     4,
					Locker: 2,
					Port:   1,
					Status: entities.LockerStatusAvailable,
				},
			},
			mockLocker: entities.Locker{
				ID: 1,
				Ports: []*entities.Port{
					{
						ID:     1,
						Locker: 1,
						Port:   1,
						Status: entities.LockerStatusAvailable,
					},
				},
			},
			expectedPort: &entities.Port{
				ID:     1,
				Locker: 1,
				Port:   1,
				Status: entities.LockerStatusAvailable,
			},
			expectedError: nil,
		},
		{
			name: "should skip to next available port if first is not in MQTT response",
			ports: []*entities.Port{
				{
					ID:     1,
					Locker: 1,
					Port:   1,
					Status: entities.LockerStatusAvailable,
				},
				{
					ID:     2,
					Locker: 1,
					Port:   2,
					Status: entities.LockerStatusAvailable,
				},
			},
			mockLocker: entities.Locker{
				ID: 1,
				Ports: []*entities.Port{
					{
						ID:     2,
						Locker: 1,
						Port:   2,
						Status: entities.LockerStatusAvailable,
					},
				},
			},
			expectedPort: &entities.Port{
				ID:     2,
				Locker: 1,
				Port:   2,
				Status: entities.LockerStatusAvailable,
			},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock repository
			mockRepo := new(MockLockerRepository)
			if tt.expectedError == nil {
				mockRepo.On("GetAvailablePorts", mock.Anything).Return(tt.mockLocker, nil)
			}

			// Create use case with mock repository
			uc := NewRegisterPackageCase(mockRepo)

			// Execute test
			result, err := uc.getAvailablePort(tt.ports)

			// Assert results
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, tt.expectedPort.ID, result.ID)
				assert.Equal(t, tt.expectedPort.Locker, result.Locker)
				assert.Equal(t, tt.expectedPort.Port, result.Port)
				assert.Equal(t, tt.expectedPort.Status, result.Status)
			}

			// Verify mock expectations
			mockRepo.AssertExpectations(t)
		})
	}
}
