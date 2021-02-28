func repeatedSubstringPattern(s string) bool {
	for i := 0; i < len(s); i++ {
		end := len(s)
		str := ""
		if i != 0 {
			end = len(s) / i
		}
		for j := 0; j < end; j++ {
			str += s[0:i]
		}
		if str == s {
			return true
		}
	}
	return false
}