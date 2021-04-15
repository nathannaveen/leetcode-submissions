func maxAscendingSum(nums []int) int {
	maximum := nums[0]
	sum := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			sum += nums[i]
		} else {
			maximum, sum = int(math.Max(float64(maximum), float64(sum))), nums[i]
		}
	}
	maximum = int(math.Max(float64(maximum), float64(sum)))
	return maximum
}