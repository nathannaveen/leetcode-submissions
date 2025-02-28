func sortArray(nums []int) []int {

	for i := 1; i < len(nums); i++ {
		if i != 0 && nums[i] < nums[i-1] {
			nums[i], nums[i - 1] = nums[i-1], nums[i]
			i -= 2
		}
	}

	return nums
}
