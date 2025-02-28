func largestPerimeter(A []int) int {
	sort.Ints(A)
	for i := len(A) - 1; i >= 2; i-- {
		if A[i] < A[i - 1] + A[i - 2] {
			return A[i] + A[i - 1] + A[i - 2]
		}
	}
	return 0
}