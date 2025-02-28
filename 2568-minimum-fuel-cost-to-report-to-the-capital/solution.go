var neighbors [][]int32
var s int32

func fuelCitiesSum(cityIdx, exceptNeighborIdx int32) (fuelSum int64, citiesSum int32) {
	citiesSum = 1
	for _, neighborIdx := range neighbors[cityIdx] {
		if neighborIdx == exceptNeighborIdx {
			continue
		}
		neighborFuelSum, neighborCitiesSum := fuelCitiesSum(neighborIdx, cityIdx)
		fuelSum += neighborFuelSum
		citiesSum += neighborCitiesSum
	}
	fuelSum += int64((citiesSum + s - 1) / s)
	return
}

func minimumFuelCost(roads [][]int, seats int) int64 {
	neighbors = make([][]int32, len(roads)+1)
	s = int32(seats)
	for _, road := range roads {
		neighbors[road[0]] = append(neighbors[road[0]], int32(road[1]))
		neighbors[road[1]] = append(neighbors[road[1]], int32(road[0]))
	}
	result, _ := fuelCitiesSum(0, 0)
	result -= int64((len(neighbors) + seats - 1) / seats)
	return result
}