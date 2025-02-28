func searchInsert(nums []int, target int) int {
	high := len(nums) - 1
	low := 0
	mid := 0
	for low <= high {
		mid = low + (high-low)/2
		if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			return mid
		}
	}
	if nums[mid] > target {
		return mid
	}
	return mid + 1
}