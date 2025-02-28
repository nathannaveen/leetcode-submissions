func secondHighest(s string) int {
	highest := -1
	second := -1

	for _, i := range s {
		if i - ':' < 0 {
			num := int(i - '0')
			if num > highest {
				second, highest = highest, num
			} else if num > second && num < highest {
				second = num
			}
		}
	}
	return second
}
