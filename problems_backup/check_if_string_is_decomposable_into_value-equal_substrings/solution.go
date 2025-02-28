func isDecomposable(s string) bool {
	twoCounter := 0
	counter := 0
	prev := "-1"

	for i := 0; i < len(s); i++ {
		if string(s[i]) == prev {
			counter++
		} else {
			done := false
			twoCounter, done = notEqualToPrev(counter, twoCounter)
			if done { return false }
			
			prev = string(s[i])
			counter = 1
		}
	}

	done := false
	twoCounter, done = notEqualToPrev(counter, twoCounter)
	if done { return false }
	
	return twoCounter == 1
}

func notEqualToPrev(counter int, twoCounter int) (int, bool) {
	if (counter-2)%3 == 0 {
		twoCounter++
	} else if counter%3 != 0 {
		return 0, true
	}
	return twoCounter, false
}