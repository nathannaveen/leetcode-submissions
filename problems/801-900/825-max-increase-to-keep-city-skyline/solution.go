func maxIncreaseKeepingSkyline(grid [][]int) int {
	leftToRight := make([]float64, len(grid))
	topToBottom := make([]float64, len(grid))
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			leftToRight[i] = math.Max(leftToRight[i], float64(grid[i][j]))
			topToBottom[j] = math.Max(topToBottom[j], float64(grid[i][j]))
		}

	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			res += int(math.Min(leftToRight[i], topToBottom[j])) - grid[i][j]
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}