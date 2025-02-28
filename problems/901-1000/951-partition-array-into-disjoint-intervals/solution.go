func partitionDisjoint(A []int) int {
	max := A[0]

	for i := 0; i < len(A); i++ {
		shouldReturn := true

		if max < A[i] {
			max = A[i]
		}
		for j := i + 1; j < len(A); j++ {
			if A[j] < max {
				shouldReturn = false
				break
			}
		}
		if shouldReturn {
			return i + 1
		}
	}

	return -1
}