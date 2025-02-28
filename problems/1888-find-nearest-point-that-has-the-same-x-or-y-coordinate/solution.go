func nearestValidPoint(x int, y int, points [][]int) int {
	minimum := 100000000 // the max value has to be 10^4 * 10^4 = 10^8 = 100,000,000
	resIndex := -1
	for i, point := range points {
		manhattanYAbs := absoluteValue(point[1] - y)
		manhattanXAbs := absoluteValue(point[0] - x)
		if point[0] == x && manhattanYAbs < minimum {
			minimum = manhattanYAbs
			resIndex = i
		} else if point[1] == y && manhattanXAbs < minimum {
			minimum = manhattanXAbs
			resIndex = i
		}
	}
	return resIndex
}

func absoluteValue(n int) int {
	if n < 0 {
		return -n
	}
	return n
}