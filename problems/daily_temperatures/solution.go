func dailyTemperatures(T []int) []int {
	h := []int{}
	for i, i2 := range T {
		containsGreater := false
		for j := i; j < len(T); j++ {
			if i2 < T[j] {
				h = append(h, j - i)
				containsGreater = true
				break
			}
		}
		if !containsGreater {
			h = append(h, 0)
		}
	}
	return h
}