func arrayRankTransform(arr []int) []int {
	g := make([]int, len(arr))
	copy(g, arr)
	m := make(map[int]int)
	sort.Ints(g)
	counter := 1
	for i := range g {
		if i >= 1 && g[i] == g[i-1] {
			counter -= 1
		}
		m[g[i]] = counter
		counter++
	}
	for i := range arr {
		arr[i] = m[arr[i]]
	}
	return arr
}
