func transpose(matrix [][]int) [][]int {
	n := len(matrix)
	m := len(matrix[0])
	res := [][]int{}

	for i := 0; i < m; i++ {
		res = append(res, []int{})
		for j := 0; j < n; j++ {
			res[i] = append(res[i], matrix[j][i])
		}
	}
	return res
}