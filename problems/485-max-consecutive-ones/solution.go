func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	counter := 0
	for _, num := range nums {
		if num == 0 {
			max = int(math.Max(float64(counter), float64(max)))
			counter = 0
		} else {
			counter++
		}
	}
	max = int(math.Max(float64(counter), float64(max)))
	return max
}