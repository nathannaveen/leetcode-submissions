func uniquePaths(m int, n int) int {
	h := make([][]int, m)
	for i := 0; i < m; i++ {
		h[i] = make([]int, n)
	}
	h[0][0] = 1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i > 0 && j > 0 {
				h[i][j] = h[i][j-1] + h[i-1][j]
			} else if i != 0 {
				h[i][j] = h[i-1][j]
			} else if j != 0 {
				h[i][j] = h[i][j-1]
			}
		}
	}
	return h[m-1][n-1]
}