func kLengthApart(nums []int, k int) bool {
	old := -1

	for i, num := range nums {
		if num == 1 {
			if old != -1 && nums[old] == 1 && i-old-1 < k {
				return false
			}
			old = i
		}
	}

	return true
}