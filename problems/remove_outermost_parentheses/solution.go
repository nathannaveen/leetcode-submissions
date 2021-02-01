func removeOuterParentheses(S string) string {
	s := ""
	h := []string{}

	for _, i2 := range S {
		if string(i2) == "(" {
			if len(h) != 0 {
				s += "("
			}
			h = append(h, "(")
		} else {
			if len(h) != 1 {
				s += ")"
			}
			h = h[:len(h) - 1]
		}
	}
	return s
}