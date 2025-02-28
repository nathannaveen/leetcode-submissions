func removePalindromeSub(s string) int {
	if s == "" {
		return 0
	}
	if palindromic(s) {
		return 1
	}
	return 2
}

func palindromic(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		if s[left] != s[right] {
			return false
		}

		left, right = left+1, right-1
	}

	return true
}
