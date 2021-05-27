func rotate(matrix [][]int) {
	for i := 0; i < len(matrix)/2; i++ {
		newI := abs(i-len(matrix)) - 1
		for j := 0; j < len(matrix); j++ {
			matrix[newI][j], matrix[i][j] = matrix[i][j], matrix[newI][j]
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := i + 1; j < len(matrix); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
