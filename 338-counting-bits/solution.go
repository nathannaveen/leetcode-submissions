func countBits(num int) []int {
	res := []int{}

	for i := 0; i <= num; i++ {
		n := i
		counter := 0
		for n > 0 {
			counter += n & 1
			n >>= 1
		}
		res = append(res, counter)
	}
	return res
}