func minOperations(logs []string) int {
	stack := []string{}

	for _, log := range logs {
		if len(stack) > 0 && log == "../" {
			stack = stack[:len(stack)-1]
		} else if log == "./" {
			continue
		} else if log != "../" {
			stack = append(stack, log)
		}
	}
	fmt.Println(stack)
	return len(stack)
}
