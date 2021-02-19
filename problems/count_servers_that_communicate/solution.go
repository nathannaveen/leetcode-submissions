func countServers(grid [][]int) int {
	counter := 0
	h := [][]int{}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				l := []int{i, j}
				h = append(h, l)
			}
		}
	}

	for i := 0; i < len(h); i++ {
		for j := 0; j < len(h); j++ {
			if h[i][0] == h[j][0] && h[i][1] != h[j][1] {
				counter++
				break
			} else if h[i][1] == h[j][1] && h[i][0] != h[j][0] {
				counter++
				break
			}
		}
	}

	return counter
}
