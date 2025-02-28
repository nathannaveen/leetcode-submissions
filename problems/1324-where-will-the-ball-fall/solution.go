func stuck(grid [][]int, j int, i int, n int) bool {
	return (j == 0 && grid[i][j] == -1) || (j == n-1 && grid[i][j] == 1) ||
		(j != 0 && grid[i][j] == -1 && grid[i][j-1] == 1) || (j != n-1 && grid[i][j] == 1 && grid[i][j+1] == -1)
}

func findBall(grid [][]int) []int {
	n := len(grid[0])
	res := make([]int, n)
	arr := make([][]int, len(grid))
	arr[0] = make([]int, n)

	for i := 0; i < n; i++ {
		if stuck(grid, i, 0, n) {
			res[i] = -1
		} else if grid[0][i] == 1 {
			arr[0][i+1] = i + 1
		} else if grid[0][i] == -1 {
			arr[0][i-1] = i + 1
		}
	}

	for i := 1; i < len(grid); i++ {
		arr[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if arr[i-1][j] != 0 {
				if stuck(grid, j, i, n) {
					res[arr[i-1][j]-1] = -1
				} else if grid[i][j] == 1 {
					arr[i][j+1] = arr[i-1][j]
				} else if grid[i][j] == -1 {
					arr[i][j-1] = arr[i-1][j]
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		if arr[len(arr)-1][i] != 0 {
			res[arr[len(arr)-1][i]-1] = i
		}
	}

	return res
}