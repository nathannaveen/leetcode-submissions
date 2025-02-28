func minInsertions(s string) int {
	res := 0
	counter := 0
	for _, i := range s {
		if i == '(' {
			counter += 2
			if counter%2 != 0 {
				res++
				counter--
			}
		} else {
			counter -= 1
			if counter < 0 {
				res++
				counter += 2
			}
		}
	}
	return res + counter
}
