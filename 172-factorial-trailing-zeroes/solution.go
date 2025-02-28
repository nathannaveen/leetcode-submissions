func trailingZeroes(n int) int {
	res := 0

	for n / 5 > 0 {
		res += n / 5
		n /= 5
	}
	
	return res
}