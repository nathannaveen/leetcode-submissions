func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	first := make([]int, 26)
	second := make([]int, 26)

	for i := 0; i < len(s); i++ {
		first[int(rune(s[i]) - 'a')] ++
		second[int(rune(t[i]) - 'a')] ++
	}

	for i := 0; i < len(first); i++ {
		if first[i] != second[i] {
			return false
		}
	}
	return true
}