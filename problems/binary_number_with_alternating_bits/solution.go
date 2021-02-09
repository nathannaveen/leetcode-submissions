func hasAlternatingBits(n int) bool {
	one := false
	if n&1 == 0 {
		one = true
	}

	for n > 0 {
		if n&1 == 0 && one == false {
			return false
		} else if n&1 == 1 && one == true {
			return false
		}
		one = !one
		n >>= 1
	}
	return true
}