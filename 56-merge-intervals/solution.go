func merge(intervals [][]int) [][]int {
	res := [][]int{}
	for i := 1; i < len(intervals); i++ {
		if i >= 1 && intervals[i-1][0] > intervals[i][0] {
			intervals[i-1], intervals[i] = intervals[i], intervals[i-1]
			i -= 2
		}
	}
	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if res[len(res)-1][1] >= intervals[i][0] {
			res[len(res)-1][1] = max(intervals[i][1], res[len(res)-1][1])
		} else {
			res = append(res, intervals[i])
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
