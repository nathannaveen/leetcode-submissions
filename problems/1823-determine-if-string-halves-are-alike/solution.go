func halvesAreAlike(s string) bool {
	first, second := make([]int, 26), make([]int, 26)
	firstCounter, secondCounter := 0, 0
	s = strings.ToLower(s)
	for i, i2 := range s {
		if i < len(s)/2 {
			first[int(i2-'a')]++
		} else {
			second[int(i2-'a')]++
		}
	}

	firstCounter += first[0] + first[4] + first[8] + first[14] + first[20]
	secondCounter += second[0] + second[4] + second[8] + second[14] + second[20]
	return firstCounter == secondCounter
}
