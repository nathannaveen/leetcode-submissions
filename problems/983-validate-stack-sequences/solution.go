func validateStackSequences(pushed []int, popped []int) bool {
	h := []int{}
	popCounter := 0
	for i := 0; i < len(pushed); i++ {
		h = append(h, pushed[i])
		for len(h) != 0 && popped[popCounter] == h[len(h)-1] {
			h = h[:len(h)-1]
			popCounter++
			if popCounter == len(popped) {
				break
			}
		}
	}

	return len(h) == 0
}
