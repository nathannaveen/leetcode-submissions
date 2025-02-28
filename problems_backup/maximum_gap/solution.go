func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	max := 0
	sort.Ints(nums)

	for i := 0; i < len(nums) - 1; i++ {
		if max < int(math.Abs(float64(nums[i]-nums[i+1]))) {
			max = int(math.Abs(float64(nums[i]-nums[i+1]))) 
		}
	}
	
	return max
}