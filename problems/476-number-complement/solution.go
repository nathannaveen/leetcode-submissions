func findComplement(num int) int {
	if num == 0 {
		return 1
	}
	res := 0
	g := 1
	for num > 0 {
		if num&1 == 0 {
			res += g
		}
		g <<= 1
		num >>= 1
	}
	return res
}