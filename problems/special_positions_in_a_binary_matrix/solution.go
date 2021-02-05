func numSpecial(mat [][]int) int {
	counter := 0
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			rowCounter := 0
			colCounter := 0
			if mat[i][j] == 1 {
				for k := 0; k < len(mat[0]); k++ {
					rowCounter += mat[i][k]
				}
				if rowCounter > 1 {
					break
				}
				for k := 0; k < len(mat); k++ {
					colCounter += mat[k][j]
				}
			}
			if rowCounter != 0 && colCounter != 0 && rowCounter <= 1 && colCounter <= 1 {
				counter ++
			}
		}
	}
	return counter
}
