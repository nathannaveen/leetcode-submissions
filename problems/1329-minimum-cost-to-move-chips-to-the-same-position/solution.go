func minCostToMoveChips(position []int) int {
	odd, even := 0, 0

	for _, i := range position {
		if i%2 == 0 {
			odd++
		} else {
			even++
		}
	}
	if even < odd {
		return even
	}
	return odd
}