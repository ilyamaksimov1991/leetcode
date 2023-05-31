package practics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParkingSystem_AddCar(t *testing.T) {
	parkingSystem := NewParkingSystem(1, 1, 0)
	assert.Equal(t, true, parkingSystem.AddCar(1))  // return true because there is 1 available slot for a big car
	assert.Equal(t, true, parkingSystem.AddCar(2))  // return true because there is 1 available slot for a medium car
	assert.Equal(t, false, parkingSystem.AddCar(3)) // return false because there is no available slot for a small car
	assert.Equal(t, false, parkingSystem.AddCar(1)) // return false because there is no available slot for a big car. It is already occupied.
}
