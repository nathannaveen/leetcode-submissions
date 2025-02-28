func restoreMatrix(rowSum []int, colSum []int) [][]int {
	res := make([][]int, len(rowSum))

	for i := 0; i < len(rowSum); i++ {
		arr := make([]int, len(colSum))

		for j := 0; j < len(colSum); j++ {
			arr[j] = colSum[j]
			if colSum[j] > rowSum[i] {
				arr[j] = rowSum[i]
			}
			rowSum[i], colSum[j] = rowSum[i] - arr[j], colSum[j] - arr[j]
		}

		res[i] = arr
	}
	return res
}