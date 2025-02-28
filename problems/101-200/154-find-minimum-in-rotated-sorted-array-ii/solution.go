func findMin(nums []int) int {
	sort.Ints(nums)
	return nums[0]
}