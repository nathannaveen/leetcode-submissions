func arrangeCoins(n int) int {
	counter := 1
	sum := 0
	for true {
		sum += counter
		if n < sum {
			return counter - 1
		} else if n == sum {
			return counter
		}
		counter += 1
	}
	return -1
}