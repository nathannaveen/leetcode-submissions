func maximum69Number(num int) int {
	newNum, product, last6Position := num, 1, 0

	for newNum > 0 {
		if newNum%10 == 6 {
			last6Position = product
		}
		newNum /= 10
		product *= 10
	}
	return num + last6Position*3
}