func sortedSquares(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		nums[i] = nums[i] * nums[i]
	}
	for i := 1; i < len(nums); i++ {
		if i >= 1 && nums[i-1] > nums[i] {
			nums[i-1], nums[i] = nums[i], nums[i-1]
			i -= 2
		}
	}
	return nums
}