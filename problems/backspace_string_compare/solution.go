func backspaceCompare(s string, t string) bool {
	for i := 0; i < len(s); i++ {
		if i > 0 && s[i] == '#' {
			s = s[:i-1] + s[i+1:]
			i -= 2
		} else if s[i] == '#' {
			s = s[1:]
			i -= 1
		}
	}
	for i := 0; i < len(t); i++ {
		if i > 0 && t[i] == '#' {
			t = t[:i-1] + t[i+1:]
			i -= 2
		} else if t[i] == '#' {
			t = t[1:]
			i -= 1
		}
	}
	return s == t
}