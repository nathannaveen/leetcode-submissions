func arraySign(nums []int) int {
	sign := 1

	for _, num := range nums {
		if num <= -1 {
			sign = - sign
		} else if num == 0 {
			return 0
		}
	}
	
	return sign
}