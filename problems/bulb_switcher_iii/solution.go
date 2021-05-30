func numTimesAllBlue(light []int) int {
	numOnLeft := 0
	res := 0

	for i := 0; i < len(light); i++ {
		if light[i] > numOnLeft {
			numOnLeft = light[i]
		}
		if numOnLeft == i + 1 {
			res++
		}
	}
	return res
}