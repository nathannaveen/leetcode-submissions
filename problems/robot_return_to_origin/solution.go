func judgeCircle(moves string) bool {
	x, y := 0, 0

	for _, i := range moves {
		if i == 'U' {
			y++
		} else if i == 'D' {
			y--
		} else if i == 'R' {
			x++
		} else if i == 'L' {
			x--
		}
	}
	return x == 0 && y == 0
}