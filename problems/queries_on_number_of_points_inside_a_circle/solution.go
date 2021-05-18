func countPoints(points [][]int, queries [][]int) []int {
	res := make([]int, len(queries))
	
	for i, query := range queries {
		for _, point := range points {
			if squared(point[0] - query[0]) + squared(point[1] - query[1]) <= squared(query[2]) {
				res[i]++
			}
		}
	}
	return res
}

func squared(a int) int {
	return a * a
}