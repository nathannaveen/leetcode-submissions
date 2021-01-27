func nextGreatestLetter(letters []byte, target byte) byte {
	h := rune(letters[0])
	for _, letter := range letters {
		if letter > target {
			h = rune(letter)
			break
		}
	}
	return byte(h)
}