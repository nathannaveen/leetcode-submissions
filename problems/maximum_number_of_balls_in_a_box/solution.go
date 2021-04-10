func countBalls(lowLimit int, highLimit int) int {
	m, maximum := make(map[int]int), 0
	for i := lowLimit; i <= highLimit; i++ {
		n, sum := i, 0

		for n > 0 {
			sum, n = sum+(n % 10), n / 10
		}
		m[sum]++
		
		if m[sum] > maximum {
			maximum = m[sum]
		}
	}
	return maximum
}