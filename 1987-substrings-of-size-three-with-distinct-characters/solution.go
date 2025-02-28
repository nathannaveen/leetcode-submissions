func countGoodSubstrings(s string) int {
	counter := 0
	res := 0
	m := make(map[uint8]int)

	for i := 0; i < int(math.Min(float64(len(s)), float64(3))); i++ {
		counter = addToCounter(s, m, i, counter, 1, 1, 2)
	}

	res = addToRes(counter, res)

	for i := 3; i < len(s); i++ {
		counter = addToCounter(s, m, i, counter, 1, 1, 2)
		
		counter = addToCounter(s, m, i - 3, counter, -1, 0, 1)

		res = addToRes(counter, res)
	}

	return res
}

func addToCounter(s string, m map[uint8]int, i int, counter, add, first, second int) int {
	m[s[i]] += add
	if m[s[i]] == first {
		counter += add
	} else if m[s[i]] == second {
		counter -= add
	}
	return counter
}

func addToRes(counter int, res int) int {
	if counter == 3 {
		res++
	}
	return res
}