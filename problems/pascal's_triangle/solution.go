func generate(numRows int) [][]int {
	res := [][]int{}

	for i := 0; i < numRows; i++ {
		//temp := make([]int, i+1)
		res = append(res, make([]int, i+1))
		res[i][0], res[i][len(res[i]) - 1] = 1, 1
		//temp[0], temp[len(temp)-1] = 1, 1
		for j := 1; j < len(res[i])-1; j++ {
			res[i][j] = res[i-1][j-1] + res[i-1][j]
		}
	}

	return res
}