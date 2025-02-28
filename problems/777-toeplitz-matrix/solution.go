func isToeplitzMatrix(matrix [][]int) bool {
	for rowCounter := 0; rowCounter < len(matrix[0]); rowCounter++ {
		for i := 0; i < len(matrix); i++ {
			col := i

			for row := rowCounter; row < len(matrix[0]); row++ {
				if col == len(matrix) {
					break
				}
				if matrix[col][row] != matrix[i][rowCounter] {
					return false
				}
				col++
			}

		}
	}
	return true
}
