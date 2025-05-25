package lockerusecases

import (
	"testing"

	"github.com/joaofilippe/pegtech/internal/domain/entities"
	"github.com/stretchr/testify/assert"
)

func TestGeneratePassword(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "should generate 6 digit numeric password",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			password, err := generatePassword()
			assert.NoError(t, err)
			assert.Len(t, password, 6)
			assert.Regexp(t, `^\d{6}$`, password)
		})
	}
}

func TestGetAvailableLocker(t *testing.T) {
	tests := []struct {
		name           string
		lockers        []*entities.Port
		expectedLocker *entities.Port
		expectedError  error
	}{
		{
			name: "should return first available locker",
			lockers: []*entities.Port{
				{ID: 1, Status: entities.LockerStatusOccupied},
				{ID: 2, Status: entities.LockerStatusAvailable},
				{ID: 3, Status: entities.LockerStatusAvailable},
			},
			expectedLocker: &entities.Port{ID: 2, Status: entities.LockerStatusAvailable},
			expectedError:  nil,
		},
		{
			name: "should return error when no available lockers",
			lockers: []*entities.Port{
				{ID: 1, Status: entities.LockerStatusOccupied},
				{ID: 2, Status: entities.LockerStatusOccupied},
			},
			expectedLocker: nil,
			expectedError:  ErrNoAvailableLockers,
		},
		{
			name:           "should return error when empty lockers list",
			lockers:        []*entities.Port{},
			expectedLocker: nil,
			expectedError:  ErrNoAvailableLockers,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			locker, err := getAvailablePort(tt.lockers)
			if tt.expectedError != nil {
				assert.ErrorIs(t, err, tt.expectedError)
				assert.Nil(t, locker)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedLocker.ID, locker.ID)
				assert.Equal(t, tt.expectedLocker.Status, locker.Status)
			}
		})
	}
}

func TestGetAvailableLockerIDs(t *testing.T) {
	tests := []struct {
		name          string
		lockers       []*entities.Port
		expectedIDs   []int
		expectedError error
	}{
		{
			name: "should return IDs of all available lockers",
			lockers: []*entities.Port{
				{ID: 1, Status: entities.LockerStatusOccupied},
				{ID: 2, Status: entities.LockerStatusAvailable},
				{ID: 3, Status: entities.LockerStatusAvailable},
			},
			expectedIDs:   []int{2, 3},
			expectedError: nil,
		},
		{
			name: "should return empty list when no available lockers",
			lockers: []*entities.Port{
				{ID: 1, Status: entities.LockerStatusOccupied},
				{ID: 2, Status: entities.LockerStatusOccupied},
			},
			expectedIDs:   []int{},
			expectedError: nil,
		},
		{
			name:          "should return empty list for empty lockers list",
			lockers:       []*entities.Port{},
			expectedIDs:   []int{},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ids, err := getAvailableLockerIDs(tt.lockers)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedIDs, ids)
		})
	}
}
