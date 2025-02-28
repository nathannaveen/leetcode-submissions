func bitwiseComplement(N int) int {
	if N == 0 {
		return 1
	}
	res := 0
	g := 1
	for N > 0 {
		if N&1 == 0 {
			res += g
		}
		g <<= 1
		N >>= 1
	}
	return res
}
