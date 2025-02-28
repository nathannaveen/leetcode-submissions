func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) < 1 || len(matrix[0]) < 1 {
		return false
	}
	i := 0
	j := len(matrix[0]) - 1
	for i < len(matrix) && j >= 0 {
		if matrix[i][j] < target {
			i++
		} else if matrix[i][j] > target {
			j--
		} else if matrix[i][j] == target {
			return true
		}
	}
	return false
}
