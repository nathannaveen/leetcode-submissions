func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func surfaceArea(grid [][]int) int {
	// 2 + shape * 4 == area of each shape
	res := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] != 0 {
				res += 2 + grid[i][j]*4
			}
			if i-1 >= 0 {
				res -= 2 * min(grid[i-1][j], grid[i][j])
			}
			if j-1 >= 0 {
				res -= 2 * min(grid[i][j-1], grid[i][j])
			}
		}
	}

	return res
}