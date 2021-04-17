func isSubsequence(s string, t string) bool {
	sCounter, i := 0, 0

	for i < len(t) && sCounter < len(s) {
		if t[i] == s[sCounter] {
			sCounter++
		}
		i++
	}
	return sCounter == len(s)
}