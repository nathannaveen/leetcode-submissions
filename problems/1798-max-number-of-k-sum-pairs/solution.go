func maxOperations(nums []int, k int) int {
	sort.Ints(nums)
	res, left, right := 0, 0, len(nums) - 1
	for left < right {
		sum := nums[left] + nums[right]
		if sum == k {
			res += 1
			left++
			right--
		} else if sum < k {
			left++
		} else {
			right--
		}
	}
	return res
}