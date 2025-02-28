func getRow(rowIndex int) []int {
	res := []int{}

	for i := 0; i <= rowIndex; i++ {
		h := make([]int, len(res))
		copy(h, res)
		res = append(res, 1)
		for j := 1; j < len(res)-1; j++ {
			res[j] = h[j-1] + h[j]
		}
	}
	return res
}