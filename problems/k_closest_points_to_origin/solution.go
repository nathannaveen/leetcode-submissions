func kClosest(points [][]int, k int) [][]int {
	for i := 1; i < len(points); i++ {
		if i >= 1 {
			distanceA := square(points[i - 1][0]) + square(points[i - 1][1])
			distanceB := square(points[i][0]) + square(points[i][1])
			if distanceA > distanceB {
				points[i - 1], points[i] = points[i], points[i - 1]
				i -= 2
			}
		}
	}
	return points[:k]
}

func square(n int) int {
	return n * n
}