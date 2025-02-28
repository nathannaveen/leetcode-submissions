func minSetSize(arr []int) int {
	h := make([]int, 100001)
	n := len(arr)
	res := 0

	for _, i := range arr {
		h[i]++
	}

	sort.Slice(h, func(i, j int) bool {
		return h[i] > h[j]
	})

	for i := 0; i < len(h) && h[i] != 0; i++ {
		n -= h[i]
		res++
		if n <= len(arr)/2 {
			return res
		}
	}
	
	return len(arr)
}