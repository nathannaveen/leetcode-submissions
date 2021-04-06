func robotSim(commands []int, obstacles [][]int) int {
	maximum := 0
	direction := 0
	x, y := 0, 0

	for _, command := range commands {
		if command == -1 { // Turn right
			direction += 1
		} else if command == -2 { // Turn left
			direction += 3
		} else {
			hitObsticle := false
			for _, obstacle := range obstacles {
				if direction%4 == 0 && x == obstacle[0] && y < obstacle[1] && y+command >= obstacle[1] { // North
					hitObsticle = true
					y = obstacle[1] - 1
				} else if direction%4 == 1 && x < obstacle[0] && y == obstacle[1] && x+command >= obstacle[0] { // East
					hitObsticle = true
					x = obstacle[0] - 1
				} else if direction%4 == 2 && x == obstacle[0] && y > obstacle[1] && y-command <= obstacle[1] { // South
					hitObsticle = true
					y = obstacle[1] + 1
				} else if direction%4 == 3 && x > obstacle[0] && y == obstacle[1] && x-command <= obstacle[0] { // West
					hitObsticle = true
					x = obstacle[0] + 1
				}
			}

			if !hitObsticle {
				if direction%4 == 0 {
					y += command
				} else if direction%4 == 1 {
					x += command
				} else if direction%4 == 2 {
					y -= command
				} else {
					x -= command
				}
			}
			maximum = max(maximum, abs(x)*abs(x)+abs(y)*abs(y))

		}

	}
	return maximum
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}