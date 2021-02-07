func findPeakElement(nums []int) int {
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			return i
		}
	}
	
	if len(nums) >= 2 && nums[len(nums)-1] > nums[len(nums)-2] {
		return len(nums) - 1
	}
	return 0
}
