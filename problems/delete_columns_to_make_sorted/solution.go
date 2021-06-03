func minDeletionSize(strs []string) int {
	res := 0
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if strs[j][i] < strs[j-1][i] {
				res++
				break
			}
		}
	}
	return res
}