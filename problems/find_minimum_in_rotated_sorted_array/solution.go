func findMin(nums []int) int {

	for i := 1; i < len(nums)-1; i++ {
		if nums[i] < nums[i-1] && nums[i] < nums[i+1] {
			return nums[i]
		}
	}

	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return int(math.Min(float64(nums[0]), float64(nums[1])))
	}
	if nums[len(nums)-1] < nums[len(nums)-2] {
		return nums[len(nums)-1]
	}
	return nums[0]
	
}