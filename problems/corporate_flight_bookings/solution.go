func corpFlightBookings(bookings [][]int, n int) []int {
	res := make([]int, n)

	for _, booking := range bookings {
		for i := booking[0] - 1; i <= booking[1] - 1; i++ {
			res[i] += booking[2]
		}
	}
	
	return res
}