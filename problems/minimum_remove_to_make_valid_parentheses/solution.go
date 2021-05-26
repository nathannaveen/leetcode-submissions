func minRemoveToMakeValid(s string) string {
	stack := []int{}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else if s[i] == ')' {
			if len(stack) != 0 {
				stack = stack[:len(stack)-1]
			} else {
				s = s[:i] + s[i+1:]
				i--
			}
		}
	}

	for len(stack) != 0 {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		s = s[:pop] + s[pop+1:]
	}
	return s
}
