func oddCells(n int, m int, indices [][]int) int {
	x := make(map[int]int)
	y := make(map[int]int)
	counter := 0
	for _, index := range indices {
		y[index[0]]++
		x[index[1]]++
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if (y[i] + x[j]) % 2 == 1 {
				counter++
			}
		}
	}
	return counter
}