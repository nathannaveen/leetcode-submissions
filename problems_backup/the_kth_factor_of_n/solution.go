func kthFactor(n int, k int) int {
	h := []int{}

	h = append(h, 1)
	for i := 2; i < n/2+1; i++ {
		if n%i == 0 {
			h = append(h, i)
		}
	}

	h = append(h, n)
	if len(h) < k {
		return -1
	}

	for i := 1; i < k+1; i++ {
		if i == k {
			return h[i-1]
		}
	}

	return -1
}