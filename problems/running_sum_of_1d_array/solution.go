func runningSum(nums []int) []int {
	h := []int{}
	
	h = append(h, nums[0])
	for i := 1; i < len(nums); i++ {
		h = append(h, h[i - 1] + nums[i])
	}
	return h;
}