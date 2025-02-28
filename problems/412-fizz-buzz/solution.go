func fizzBuzz(n int) []string {
	h := []string{}

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			h = append(h, "FizzBuzz")
		} else if i%3 == 0 {
			h = append(h, "Fizz")
		} else if i%5 == 0 {
			h = append(h, "Buzz")
		} else {
			h = append(h, strconv.Itoa(i))
		}
	}
	return h
}
