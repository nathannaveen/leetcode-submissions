func sortTransformedArray(nums []int, a int, b int, c int) []int {
	for i := 0; i < len(nums); i++ {
		nums[i] = a*nums[i]*nums[i] + b*nums[i] + c
	}

	for i := 1; i < len(nums); i++ {
		if i >= 1 && nums[i] < nums[i-1] {
			nums[i], nums[i-1] = nums[i-1], nums[i]
			i -= 2
		}
	}

	return nums
}