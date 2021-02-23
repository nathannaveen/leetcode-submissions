func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	low := 0
	high := len(matrix) * len(matrix[0]) -1

	for low <= high {
		mid := (low + high)/2
		i := mid / len(matrix[0])
		j := mid % len(matrix[0])

		h := matrix[i][j]

		if h == target {
			return true
		}

		if h > target {
			high = mid-1
		} else {
			low = mid+1
		}
	}

	return false
}
