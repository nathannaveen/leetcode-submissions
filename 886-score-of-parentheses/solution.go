func scoreOfParentheses(S string) int {
	g := 0
	h := []int{}

	for _, i2 := range S {
		if string(i2) == "(" {
			h, g = append(h, g), 0
		} else {
			x := 0
			x, h = h[len(h)-1], h[:len(h)-1]
			g = int(math.Max(float64(2*g), 1)) + x
		}
	}
	return g
}