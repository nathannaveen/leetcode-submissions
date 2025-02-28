func longestOnes(A []int, K int) int {
	max := 0

	for i := 0; i < len(A); i++ {
		temp := 0
		counter := 0
		for j := i; j < len(A); j++ {
			if counter == K {
				for k := j; k < len(A); k++ {
					if A[k] == 0 {
						break
					}
					temp++
				}
				break
			}

			if A[j] == 0 {
				counter++
			}
			temp++
		}
		max = int(math.Max(float64(temp), float64(max)))
		if counter < K {
			break
		}
	}

	return max
}