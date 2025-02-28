func setZeroes(matrix [][]int) {
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			if matrix[row][col] != 0 {
				continue
			}
			for row2 := 0; row2 < len(matrix); row2++ {
				if matrix[row2][col] != 0 {
					matrix[row2][col] = 2147483649
				}
			}
			for col2 := 0; col2 < len(matrix[0]); col2++ {
				if matrix[row][col2] != 0 {
					matrix[row][col2] = 2147483649
				}
			}
		}
	}
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			if matrix[row][col] == 2147483649 {
				matrix[row][col] = 0
			}
		}
	}
}