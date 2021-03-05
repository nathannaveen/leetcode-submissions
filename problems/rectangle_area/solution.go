func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	sum := (D - B) * (C - A) + (H - F) * (G - E)

	left, right := max(A, E), min(G, C)
	up, down := min(D, H), max(F, B)

	if right > left && up > down {
		sum -= (right - left) * (up - down) // overlap
	}
	return sum
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
