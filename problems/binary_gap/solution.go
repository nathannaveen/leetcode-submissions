func binaryGap(n int) int {
	h := 0
	res := 0

	if n & (n-1) == 0 {
		return 0
	}

	for ; n > 0; n >>= 1 {
		if n&1 == 1 {
			if res == 0 {
				res = 1
			}
			if h > res {
				res = h
			}
			h = 0
		}
		if res > 0 {
			h++
		}
	}
	return res
}