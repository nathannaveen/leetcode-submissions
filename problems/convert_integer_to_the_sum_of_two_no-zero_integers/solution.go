
func getNoZeroIntegers(n int) []int {
	for i := 1; i < n; i++ {
		containsZero := false
		l := i
		k := n - i
		for l > 0 {
			if l % 10 == 0 {
				containsZero = true
				break
			}
			l /= 10
		}
		if !containsZero {
			for k > 0 {
				if k % 10 == 0 {
					containsZero = true
					break
				}
				k /= 10
			}
		}
		if !containsZero {
			return []int{i, n - i}
		}
	}
	return []int{}
}