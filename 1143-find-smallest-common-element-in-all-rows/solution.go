func smallestCommonElement(mat [][]int) int {
	m := make(map[int]int)

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			m[mat[i][j]]++
			if m[mat[i][j]] == len(mat) {
				return mat[i][j]
			}
		}
	}
	return -1
}
