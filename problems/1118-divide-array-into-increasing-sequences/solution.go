func canDivideIntoSubsequences(nums []int, K int) bool {
	counter, maximum := 0, 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] > nums[i] {
			maximum = max(maximum, counter)
			counter = 1
		} else {
			counter++
		}
	}
    maximum = max(maximum, counter)
	return len(nums) >= K*maximum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}