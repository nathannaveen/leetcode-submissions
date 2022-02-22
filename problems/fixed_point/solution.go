func fixedPoint(arr []int) int {
	for i := range arr {
		if i == arr[i] {
			return arr[i]
		}
	}
	return -1
}