func numKLenSubstrNoRepeats(S string, K int) int {
	if K > len(S) {
		return 0
	}
	counter, res, m := 0, 0, make(map[uint8]int)

	for i := 0; i < K; i++ {
		m[S[i]]++
		if m[S[i]] == 2 {
			counter++
		}
	}
	if counter == 0 {
		res++
	}

	for i := K; i < len(S); i++ {
		m[S[i-K]]--

		if m[S[i-K]] == 1 {
			counter--
		}
		m[S[i]]++

		if m[S[i]] == 2 {
			counter++
		}
		if counter == 0 {
			res++
		}
	}
	return res
}