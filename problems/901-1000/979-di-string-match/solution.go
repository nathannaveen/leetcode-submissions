func diStringMatch(S string) []int {
	upper, lower := len(S), 0
	res := make([]int, len(S) + 1)
	for i, i2 := range S {
		if i2 == 'I' {
			res[i] = lower
			lower++
		} else {
			res[i] = upper
			upper--
		}
	}
	res[len(S)] = lower
	return res
}