func largestUniqueNumber(A []int) int {
	h := make([]int, 1001)

	for _, i := range A {
		h[i]++
	}
	for i := 1000; i >= 0; i-- {
		if h[i] == 1 {
			return i
		}
	}
	return -1
}