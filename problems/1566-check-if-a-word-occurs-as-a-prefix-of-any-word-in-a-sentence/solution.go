func isPrefixOfWord(sentence string, searchWord string) int {
	arr := strings.Split(sentence, " ")

	for i, i2 := range arr {
		counter := 0
		for j := 0; j < len(i2); j++ {
			if j == len(searchWord) {
				break
			} else if string(searchWord[j]) == string(i2[j]) {
				counter++
			}
		}
		if counter == len(searchWord) {
			return i + 1
		}
	}

	return -1
}