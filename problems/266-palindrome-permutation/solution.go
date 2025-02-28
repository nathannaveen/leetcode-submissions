func canPermutePalindrome(s string) bool {
	m := make(map[rune]int)
	numberOfOnes := 0
	for _, i := range s {
		m[i]++
	}

	for _, i := range m {
		numberOfOnes += i % 2
	}

	return numberOfOnes <= 1
}