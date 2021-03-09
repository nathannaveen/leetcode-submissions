func minimumLength(s string) int {

	for len(s) > 1 && s[0] == s[len(s)-1] {
		common := s[:1]
		lenS := len(s)

		for i := 0; i < lenS; i++ {
			if s[len(s)-1:] == common {
				s = s[:len(s)-1]
			} else {
				break
			}
		}
		lenS = len(s)
		for i := 0; i < lenS; i++ {
			if s[:1] == common {
				s = s[1:]
			} else {
				break
			}
		}

	}

	return len(s)
}