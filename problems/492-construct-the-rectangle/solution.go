func constructRectangle(area int) []int {
	if area == 1 {
		return []int{1, 1}
	}
	res := []int{10000001, 1}
	for i := 1; i < area/2+1; i++ {
		if area%i == 0 {
			l := i
			w := area / i
			if area/i > i {
				l = area/i
				w = i
			}
			if l-w < res[0]-res[1] {
				res[0] = l
				res[1] = w
			}
		}
	}
	return res
}