func maxRotateFunction(A []int) int {
	// if len(A) == 0 || len(A) == 1 {
	// 	return 0
	// }
	max := 0

	for i := 0; i < len(A); i++ {
		sum := 0
		cur := A[len(A)-1]
		for i := len(A) - 2; i >= 0; i-- {
			A[i+1] = A[i]
			sum += (i + 1) * A[i]
		}
		A[0] = cur
		if i == 0 {
			max = sum
		} else{
			max = int(math.Max(float64(sum), float64(max)))
		}
		
	}

	return max
}