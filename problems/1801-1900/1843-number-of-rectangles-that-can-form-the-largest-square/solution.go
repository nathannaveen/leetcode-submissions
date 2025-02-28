func countGoodRectangles(rectangles [][]int) int {
	max := 0
	counter := 0
	for _, rectangle := range rectangles {
		min := int(math.Min(float64(rectangle[0]), float64(rectangle[1])))
		if min > max {
			counter = 1
			max = min
		} else if min == max {
			counter ++
		}
	}
	return counter
}