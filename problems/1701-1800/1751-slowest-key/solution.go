type slowest struct {
	maximum int
	letter  byte
}

func slowestKey(releaseTimes []int, keysPressed string) byte {
	maxAndLetter := slowest{maximum: releaseTimes[0], letter: keysPressed[0]}

	for i := 1; i < len(releaseTimes); i++ {
		difference := releaseTimes[i] - releaseTimes[i-1]
		if difference > maxAndLetter.maximum {
			maxAndLetter.maximum = difference
			maxAndLetter.letter = keysPressed[i]
		} else if difference == maxAndLetter.maximum && keysPressed[i] > maxAndLetter.letter {
			maxAndLetter.maximum = difference
			maxAndLetter.letter = keysPressed[i]
		}
	}

	return maxAndLetter.letter
}