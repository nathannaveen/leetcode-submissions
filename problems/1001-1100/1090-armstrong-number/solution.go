func isArmstrong(n int) bool {
	lengthOfNumber := len(strconv.Itoa(n))
	sum := 0
	newN := n

	for newN > 0 {
		sum += powerOfK(lengthOfNumber, newN%10)
		newN /= 10
	}
	return sum == n
}

func powerOfK(lengthOfNumber, n int) int {
	res := 1
	for i := 0; i < lengthOfNumber; i++ {
		res *= n
	}
	return res
}
