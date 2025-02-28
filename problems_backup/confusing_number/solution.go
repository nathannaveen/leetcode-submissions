func confusingNumber(N int) bool {
	newNumber := 0
	newN := N
	for N > 0 {
		lastDigit := N % 10
		newNumber *= 10
		switch lastDigit {
		case 0:
			newNumber += 0
		case 1:
			newNumber += 1
		case 6:
			newNumber += 9
		case 8:
			newNumber += 8
		case 9:
			newNumber += 6
		default:
			return false
		}
		
		N /= 10
	}
	return newNumber != newN
}