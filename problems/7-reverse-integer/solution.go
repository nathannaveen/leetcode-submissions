func reverse(x int) int {
	num := 0
	n := x

	if x < 0 {
		n *= -1
	}

	for n > 0 {
		num *= 10
		if x < 0 {
			num -= n % 10
		} else {
			num += n % 10
		}
		n /= 10
	}

	if num >= 2147483648 || num <= -2147483648 { // max and min values
		return 0
	}
	
	return num
}