func removeVowels(s string) string {
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' || s[i] == 'e' || s[i] == 'i' || s[i] == 'o' || s[i] == 'u' {
			s = s[:i] + s[i+1:]
			i -= 1
		}
	}
	return s
}