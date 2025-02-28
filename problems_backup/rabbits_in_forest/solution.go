func numRabbits(answers []int) int {
	h := make(map[int]int)
	num := 0

	for _, number := range answers {
		if number == 0 {
			num += 1
			continue
		}
		h[number]++

		if h[number] == number+1 {
			num += number + 1
			delete(h, number)
		}
	}

	for i := range h {
		num += i + 1
	}

	return num
}