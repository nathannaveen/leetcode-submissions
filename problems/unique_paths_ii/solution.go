func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	h := make([][]int, len(obstacleGrid))
	for i := 0; i < len(obstacleGrid); i++ {
		h[i] = make([]int, len(obstacleGrid[0]))
	}

	h[0][0] = 1

	for i := 0; i < len(obstacleGrid); i++ {
		for j := 0; j < len(obstacleGrid[0]); j++ {
			if obstacleGrid[i][j] == 1 {
				h[i][j] = 0
			} else if i > 0 && j > 0 {
				h[i][j] = h[i][j-1] + h[i-1][j]
			} else if i != 0 {
				h[i][j] = h[i-1][j]
			} else if j != 0 {
				h[i][j] = h[i][j-1]
			}
		}
	}

	return h[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}
