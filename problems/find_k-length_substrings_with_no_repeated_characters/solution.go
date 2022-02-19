func numKLenSubstrNoRepeats(s string, k int) int {
    if k > len(s) {
		return 0
	}
	counter, res, m := 0, 0, make(map[uint8]int)

	for i := 0; i < k; i++ {
		m[s[i]]++
		if m[s[i]] == 2 {
			counter++
		}
	}
	if counter == 0 {
		res++
	}

	for i := k; i < len(s); i++ {
		m[s[i-k]]--

		if m[s[i-k]] == 1 {
			counter--
		}
		m[s[i]]++

		if m[s[i]] == 2 {
			counter++
		}
		if counter == 0 {
			res++
		}
	}
	return res
}
