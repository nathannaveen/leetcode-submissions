func compress(chars []byte) int {
	resCounter := 0
	counter := 0
	letter := byte(' ')

	for _, char := range chars {
		if char != letter {
			resCounter = foo(chars, counter, letter, resCounter)
			letter = char
			counter = 1
		} else {
			counter++
		}
	}
	resCounter = foo(chars, counter, letter, resCounter)
	return resCounter
}

func foo(chars []byte, counter int, letter byte, resCounter int) int {
	if counter == 1 {
		chars[resCounter] = letter
		resCounter += 1
	} else if counter != 0 {
		chars[resCounter] = letter
		i := 0
		resCounter += 1
		h := strconv.Itoa(counter)
		for i < len(h) {
			chars[resCounter] = h[i]
			resCounter += 1
			i++
		}
	}
	return resCounter
}