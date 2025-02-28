func reorderedPowerOf2(N int) bool {
	h := strings.Split(strconv.Itoa(N), "")
	product := 1
	g := N
	for g > 0 {
		if g%10 != 0 {
			product *= g % 10
		}
		g /= 10
	}
	n := 1
	fmt.Println(h)
	for n <= 1000000000 {
		s := strconv.Itoa(n)
		notContains := false
		l := n
		smallProduct := 1
		for i := 0; i < len(h); i++ {
			if len(s) != len(h) {
				break
			}
			notContains = true
			if l%10 != 0 {
				smallProduct *= l % 10
			}
			l /= 10
			if !contains(h, string(s[i])) {
				notContains = false
				break
			}
		}
		if notContains && smallProduct == product {
			return true
		}
		n *= 2
		fmt.Println(n)
	}

	return false
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
