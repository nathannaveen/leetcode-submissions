func executeInstructions(n int, startPos []int, s string) []int {
	res := make([]int, len(s))

	for k := 0; k < len(s); k++ {
		pos := []int{ startPos[0], startPos[1] }
		S := s[k:]
		outOfBounds := false

		for i := 0; i < len(S); i++ {
			switch S[i] {
			case 'R':
				pos[1]++
			case 'D':
				pos[0]++
			case 'L':
				pos[1]--
			case 'U':
				pos[0]--
			}

			if pos[0] == -1 || pos[1] == -1 || pos[0] == n || pos[1] == n {
				outOfBounds = true
				res[k] = i
				break
			}
		}

		if !outOfBounds { res[k] = len(S) }
	}

	return res
}