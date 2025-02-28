func minStartValue(nums []int) int {
	res := 0
	sum := 0
	for _, num := range nums {
		sum += num
		res = int(math.Min(float64(res), float64(sum)))
	}
	return 1 - res
}