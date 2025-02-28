func countSubstrings(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		for j := 0; j <= len(s)-i; j++ {
			if s[j : j+i] != "" && isPalindromic(s[j : j+i]) {
				res++
			}
		}
	}
	if isPalindromic(s) {
		res++
	}
	return res
}

func isPalindromic(s string) bool {
	left, right := 0, len(s)-1

	for left <= right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}