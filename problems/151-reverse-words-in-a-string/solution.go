func reverseWords(s string) string {
	res := ""
	stack := strings.Split(s, " ")

	for _, s2 := range stack {
		if s2 != "" {
			res = s2 + " " + res
		}
	}
	return strings.Trim(res, " ")
}