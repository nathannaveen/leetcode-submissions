func longestPalindrome(s string) string {
	max := ""
	for i := 0; i < len(s); i++ {
		for j := 0; j <= len(s)-i; j++ {
			if isPalindromic(s[j : j + i]) {
				max = s[j : j + i]
			}
		}
	}
	if isPalindromic(s) {
		max = s
	}
	return max
}

func isPalindromic(s string) bool {
	left, right := 0, len(s) - 1

	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
