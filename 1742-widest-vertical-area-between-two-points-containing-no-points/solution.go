func maxWidthOfVerticalArea(points [][]int) int {
	h := []int{}
	for _, point := range points {
		h = append(h, point[0])
	}
	sort.Ints(h)
	max := 0

	for i := 0; i < len(h) - 1; i++ {
		if h[i + 1] - h[i] > max {
			max = h[i + 1] - h[i]
		}
	}
	
	return max
}