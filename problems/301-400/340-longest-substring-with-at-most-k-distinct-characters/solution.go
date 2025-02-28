func lengthOfLongestSubstringKDistinct(s string, k int) int {
	m := make(map[byte]int)
	maximum := 0
	begining := 0

	for i := 0; i < len(s); i++ {
		m[s[i]]++
		for len(m) > k {
			m[s[begining]]--
			if m[s[begining]] == 0 {
				delete(m, s[begining])
			}
			begining++
		}
		if i-begining+1 > maximum {
			maximum = i - begining + 1
		}
	}

	return maximum
}