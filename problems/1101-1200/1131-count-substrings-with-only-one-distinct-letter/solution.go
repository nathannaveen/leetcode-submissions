func countLetters(S string) int {
	res := 0
	letter := ""
	counter := 0

	for _, i := range S {
		if letter != string(i) {
			res += (counter * (counter + 1)) / 2
			counter = 0
			letter = string(i)
		}
		counter++
	}
	res += (counter * (counter + 1)) / 2
	return res
}