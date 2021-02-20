func shiftingLetters(S string, shifts []int) string {
	letters := []rune{}
	for _, i := range S {
		letters = append(letters, i)
	}

	for i := 0; i < len(S); i++ {
		for j := 0; j <= i; j++ {
			newLetterInt := int(letters[j]) + shifts[i]
			if newLetterInt > 122 { // overflow of alphabets
				newLetterInt = (newLetterInt-123)%26 + 'a'
			}
			letters[j] = rune(newLetterInt)
		}
	}
	return string(letters)
}
