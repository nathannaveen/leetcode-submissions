func sumOfDigits(A []int) int {
	minimum := A[0]
	sum := 0

	for _, i := range A {
		if i < minimum {
			minimum = i
		}
	}

	for minimum > 0 {
		sum += minimum % 10
		minimum /= 10
	}
	if sum%2 == 0 {
		return 1
	}
	return 0
}