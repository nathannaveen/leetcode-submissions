func isRobotBounded(instructions string) bool {
	x, y, degree := 0, 0, 0
	addSubtractXAndY := []int{1, 1, -1, -1}

	for _, i := range instructions {
		if i == 'R' {
			degree = (degree + 1) % 4
		} else if i == 'L' {
			degree = (degree + 3) % 4
		} else {
			if degree == 0 || degree == 2 {
				y += addSubtractXAndY[degree]
			} else {
				x += addSubtractXAndY[degree]
			}
		}
	}
	return x == 0 && y == 0 || degree != 0
}
