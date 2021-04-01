func fixedPoint(arr []int) int {
	for i, n := range arr {
		if i == n {
			return n
		}
	}
	return -1
}