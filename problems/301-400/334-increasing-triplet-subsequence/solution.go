func increasingTriplet(nums []int) bool {
	g := 2147483648
	h := 2147483648

	for _, num := range nums {
		if num <= g {
			g = num
		} else if num <= h {
			h = num
		} else {
			return true
		}
	}
	return false
}