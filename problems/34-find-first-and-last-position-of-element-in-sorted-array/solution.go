func searchRange(nums []int, target int) []int {
	left := 0
	right := len(nums) - 1

	for (left < right) && (nums[left] != target || nums[right] != target) {
		if nums[left] != target {
			left++
		}
		if nums[right] != target {
			right--
		}
	}

	if len(nums) > 0 && nums[left] == target && nums[right] == target {
		return []int{left, right}
	}

	return []int{-1, -1}
}
