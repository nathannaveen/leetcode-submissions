func maxDepthAfterSplit(seq string) []int {
	res := make([]int, len(seq))
	pos := 0
	for i, i2 := range seq {
		if i2 == '(' {
			pos++
			res[i] = pos % 2
		} else {
			res[i] = pos % 2
			pos--
		}
	}
	return res
}