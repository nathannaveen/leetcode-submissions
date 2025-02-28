func removeInterval(intervals [][]int, toBeRemoved []int) [][]int {
	res := [][]int{}
	for _, interval := range intervals {
		if interval[0] < toBeRemoved[0] && toBeRemoved[1] < interval[1] {
			res = append(res, []int{interval[0], toBeRemoved[0]}, []int{toBeRemoved[1], interval[1]})
		} else if interval[0] < toBeRemoved[0] && interval[1] > toBeRemoved[0] && interval[1] <= toBeRemoved[1] {
			res = append(res, []int{interval[0], toBeRemoved[0]})
		} else if interval[0] >= toBeRemoved[0] && interval[0] < toBeRemoved[1] && interval[1] > toBeRemoved[1] {
			res = append(res, []int{toBeRemoved[1], interval[1]})
		} else if interval[0] >= toBeRemoved[0] && interval[0] < toBeRemoved[1] && interval[1] <= toBeRemoved[1] && interval[1] > toBeRemoved[0] {
			continue
		} else {
			res = append(res, interval)
		}
	}
	return res
}