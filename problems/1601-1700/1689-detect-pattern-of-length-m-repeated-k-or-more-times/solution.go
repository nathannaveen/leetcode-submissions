func containsPattern(arr []int, m int, k int) bool {
	counter := 0

	for i := 0; i < len(arr) - m; i++ {
		if arr[i] != arr[i+m] {
			counter = 0
		} else {
			counter++
		}

		if counter == (k-1)*m {
			return true
		}
	}

	return false
}
