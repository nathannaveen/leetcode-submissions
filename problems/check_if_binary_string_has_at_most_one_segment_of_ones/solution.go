func checkOnesSegment(s string) bool {

	for i := 0; i < len(s)-1; i++ {
		if s[i:i+2] == "01" {
			return false
		}
	}
	return true
}