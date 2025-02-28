func canAttendMeetings(intervals [][]int) bool {
	for i := 1; i < len(intervals); i++ {
		if i >= 1 && intervals[i-1][0] > intervals[i][0] {
			intervals[i-1], intervals[i] = intervals[i], intervals[i-1]
			i -= 2
		}
		if i >= 1 && intervals[i-1][0] <= intervals[i][0] && intervals[i - 1][1] > intervals[i][0] {
			return false
		}
	}
	return true
}