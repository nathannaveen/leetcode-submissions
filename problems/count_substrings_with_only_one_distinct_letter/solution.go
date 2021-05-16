func countLetters(S string) int {
	res := 0
	letter := ""
	counter := 0

	for _, i := range S {
		if letter == string(i) {
			counter++
		} else {
			res += counter * (counter + 1)
			counter = 1
			letter = string(i)
		}
	}
	res += counter * (counter + 1)
	return res / 2
}
