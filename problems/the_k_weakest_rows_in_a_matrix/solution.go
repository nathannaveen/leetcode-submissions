func kWeakestRows(mat [][]int, k int) []int {
	h := make([][]int, len(mat))
	res := []int{}
	for i := 0; i < len(mat); i++ {
		counter := 0
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 0 {
				break
			}
			counter++
		}
		h[i] = append(h[i], counter, i)
	}

	for i := 1; i < len(h); i++ {
		if h[i][0] <= h[i-1][0] {
			g := h[i]
			for j := i - 1; j >= 0; j-- {
				if h[j][0] < g[0] {
					h[j+1] = g
					break
				} else if h[j][0] == g[0] && h[j][1] < g[1] {
					h[j+1] = g
					break
				}

				h[j+1] = h[j]
				if j == 0 {
					h[0] = g
				}
			}
		}
	}
	for i := 0; i < k; i++ {
		res = append(res, h[i][1])
	}
	return res
}
