func projectionArea(grid [][]int) int {
	result := 0

	for i := 0; i < len(grid); i++ {
		maxRow, maxCol := 0, 0

		for j := 0; j < len(grid[i]); j++ {
			if grid[j][i] > maxRow {
				maxRow = grid[j][i]
			}
			if grid[i][j] > maxCol {
				maxCol = grid[i][j]
			}
			if grid[i][j] != 0 {
				result++
			}
		}
		result = result + maxCol + maxRow
	}

	return result
}