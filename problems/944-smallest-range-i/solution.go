func smallestRangeI(A []int, K int) int {
	max, min := A[0], A[0]

	for _, i := range A {
		if i > max {max = i}
		if i < min {min = i}
	}
	if max-min-2*K > 0 {return max - min - 2*K}
	return 0
}
