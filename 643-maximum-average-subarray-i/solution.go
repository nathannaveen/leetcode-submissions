func findMaxAverage(nums []int, k int) float64 {

	sum := 0

	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	max := float64(sum) / float64(k)
	
	for i := k; i < len(nums); i++ {
		sum -= nums[i-k]
		sum += nums[i]
		if float64(sum) / float64(k) > max {
			max = float64(sum) / float64(k)
		}
	}
	return max
}