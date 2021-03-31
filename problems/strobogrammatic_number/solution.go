func isStrobogrammatic(num string) bool {
	newNumber := ""
	for i := len(num) - 1; i >= 0; i-- {
		lastDigit := num[i]
		switch lastDigit {
		case '0':
			newNumber += "0"
		case '1':
			newNumber += "1"
		case '6':
			newNumber += "9"
		case '8':
			newNumber += "8"
		case '9':
			newNumber += "6"
		default:
			return false
		}
	}
	return newNumber == num
}