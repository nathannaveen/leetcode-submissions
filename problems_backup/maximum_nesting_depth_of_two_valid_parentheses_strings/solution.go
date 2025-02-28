func maxDepthAfterSplit(seq string) []int {
	res := make([]int, len(seq))
	evenOrOdd := 0
	for i, i2 := range seq {
		if i2 == '(' {
			evenOrOdd++
			res[i] = evenOrOdd % 2
		} else {
			res[i] = evenOrOdd % 2
			evenOrOdd--
		}
	}
	return res
}

