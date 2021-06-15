type ParkingSystem struct {
	cars []int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	c := ParkingSystem{
		[]int{big, medium, small},
	}
	return c
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.cars[carType - 1] >= 1 {
		this.cars[carType - 1]--
		return true
	}
	return false
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */