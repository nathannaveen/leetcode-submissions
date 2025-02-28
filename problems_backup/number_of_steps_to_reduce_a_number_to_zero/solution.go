func numberOfSteps(num int) int {
	counter := 0

	for num > 0 {
		if num%2 == 1 {
			counter++
			num -= 1
		} else {
			counter++
			num /= 2	
		}
		
	}

	return counter
}
