func lengthOfLongestSubstring(s string) int {
	maximum := 0
	start := 0
	m := make(map[byte]int)
	counter := 0
	for i := 0; i < len(s); i++ {
		m[s[i]]++
		counter++

		if m[s[i]] == 2 {
			maximum = max(maximum, counter-1)
			for j := start; j < i; j++ {
				m[s[j]]--
				start++
				counter--
				if s[j] == s[i] {
					break
				}
			}
		}
	}
	maximum = max(maximum, counter)

	return maximum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}