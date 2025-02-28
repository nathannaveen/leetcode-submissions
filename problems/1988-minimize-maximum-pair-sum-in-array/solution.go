func minPairSum(nums []int) int {
	sort.Ints(nums)
	res := 0
	for i := 0; i < len(nums)/2; i++ {
		if nums[i]+nums[len(nums)-i-1] > res {
			res = nums[i] + nums[len(nums)-i-1]
		}
	}
	return res
}
