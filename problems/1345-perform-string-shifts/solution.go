func stringShift(s string, shift [][]int) string {
	counter := 0

	for _, ints := range shift {
		if ints[0] == 0 {
			counter -= ints[1]
		} else {
			counter += ints[1]
		}
	}

	for counter != 0 {
		if counter < 0 {
			s = s[1:] + string(s[0])
			counter++
		} else {
			s = string(s[len(s)-1]) + s[:len(s)-1]
			counter--
		}
	}
	return s
}