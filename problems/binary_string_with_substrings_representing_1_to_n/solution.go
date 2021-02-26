func queryString(S string, N int) bool {
	for i := 1; i <= N; i++ {
		s := ""
		n := i

		for n > 0 {
			if n&1 == 1 {
				s = "1" + s
			} else {
				s = "0" + s
			}
			n >>= 1
		}
		if !strings.Contains(S, s) {
			return false
		}
	}
	return true
}