func subtractProductAndSum(n int) int {
	product := 1
	sum := 0
	for n > 0 {
		sum += n % 10
		product *= n % 10
		n /= 10
	}
	return product - sum
}