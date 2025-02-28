func largestAltitude(gain []int) int {
	max := 0
	counter := 0

	for _, i := range gain{
		counter += i
		max = int(math.Max(float64(counter), float64(max)))
	}

	return max
}
