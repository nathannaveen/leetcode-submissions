func minOperations(boxes string) []int {
	res := make([]int, len(boxes))

	for i := 0; i < len(boxes); i++ {
		for j := 0; j < len(boxes); j++ {
			if boxes[j] == '1' && j != i {
				res[i] += abs(i - j)
			}
		}
	}
	return res
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}