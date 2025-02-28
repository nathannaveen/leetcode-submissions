func longestConsecutive(nums []int) int {
	sort.Ints(nums)
	max := 0
	num := -10
	counter := 0
	for _, n := range nums {
		if n == num {
			continue
		} else if n != num+1 {
			max = int(math.Max(float64(counter), float64(max)))
			counter = 1
			num = n
		} else {
			counter++
			num = n
		}
	}
	max = int(math.Max(float64(counter), float64(max)))
	return max
}