func maxScore(s string) int {
	maximum := 0
	sum := 0

	for _, i := range s {
		sum += int(i) - 48
	}

	sum = addAndSubtractFromOnesAndZeros2(s, sum, 0)
	maximum = sum

	for i := 1; i < len(s)-1; i++ {
		sum = addAndSubtractFromOnesAndZeros2(s, sum, i)
		if sum > maximum {
			maximum = sum
		}
	}

	return maximum
}

func addAndSubtractFromOnesAndZeros2(s string, sum, i int) int {
	if s[i] == '0' {
		return sum + 1
	}
	return sum - 1
}