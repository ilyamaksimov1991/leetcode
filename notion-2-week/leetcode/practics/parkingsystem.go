package practics

import "sync"

type carType int

const (
	bigCarType    = 1
	mediumCarType = 2
	smallCarType  = 3
)

type ParkingSystem struct {
	mu             sync.Mutex
	carTypeToCount map[carType]int
}

func NewParkingSystem(big int, medium int, small int) ParkingSystem {
	carTypeToCount := make(map[carType]int, 0)
	carTypeToCount[bigCarType] = big
	carTypeToCount[mediumCarType] = medium
	carTypeToCount[smallCarType] = small
	return ParkingSystem{
		carTypeToCount: carTypeToCount,
	}
}

func (ps *ParkingSystem) AddCar(ct int) bool {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	if res, ok := ps.carTypeToCount[carType(ct)]; ok && res > 0 {
		ps.carTypeToCount[carType(ct)]--
		return true
	}

	return false
}
