func hammingDistance(x int, y int) int {
	res := 0
	for x > 0 || y > 0 {
		if x&1 != y&1 {
			res++
		}
		x >>= 1
		y >>= 1
	}
	return res
}