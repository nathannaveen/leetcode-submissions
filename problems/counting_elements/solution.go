func countElements(arr []int) int {
	m := make(map[int]int)
	res := 0

	for _, i := range arr {
		m[i]++
	}

	for _, i := range arr {
		if m[i + 1] > 0 {
			res++
		}
	}
	return res
}