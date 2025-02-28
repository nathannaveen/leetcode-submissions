func asteroidCollision(asteroids []int) []int {
	stack := []int{}
	stack = append(stack, asteroids[0])
	for i := 1; i < len(asteroids); i++ {
		if asteroids[i] < 0 {
			g := true
			for len(stack) > 0 {
				if stack[len(stack)-1] < int(math.Abs(float64(asteroids[i]))) && stack[len(stack)-1] > 0 {
					stack = stack[:len(stack)-1]
					g = true
				} else if stack[len(stack)-1] == int(math.Abs(float64(asteroids[i]))) && stack[len(stack)-1] > 0 {
					stack = stack[:len(stack)-1]
					g = false
					break
				} else if stack[len(stack)-1] < 0 && asteroids[i] < 0 {
					g = true
					break
				} else {
					g = false
					break
				}
			}
			if g {
				stack = append(stack, asteroids[i])
			}
		} else {
			stack = append(stack, asteroids[i])
		}
	}

	return stack
}