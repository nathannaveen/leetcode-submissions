func hasAllCodes(s string, k int) bool {
	m := make(map[string]int)

	for i := 0; i < len(s)-k+1; i++ {
		m[s[i:i+k]]++
	}

	return len(m) == 1 << k
}
