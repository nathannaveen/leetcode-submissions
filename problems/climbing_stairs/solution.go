func climbStairs(n int) int {
	if n <= 1 {
		return n
	}
	total := 0
	old := 1
	oldold := 1
	for i := 2; i <= n; i++ {
		total = old + oldold
		oldold = old
		old = total
	}
	return total
}