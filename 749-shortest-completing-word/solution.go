func shortestCompletingWord(licensePlate string, words []string) string {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43,
		47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}
	product := 1
	product = productOfWord(licensePlate, product, primes)
	shortest, length := "", 16

	for _, word := range words {
		if productOfWord(word, 1, primes ) % product == 0 && len(word) < length {
			shortest, length = word, len(word)
		}
	}
	return shortest
}

func productOfWord(s string, product int, primes []int) int {
	for _, i := range s {
		lowerIndex, higherIndex := i - 'a', i - 'A'
		if lowerIndex >= 0 && lowerIndex < 26 {
			product *= primes[lowerIndex]
		} else if higherIndex >= 0 && higherIndex < 26 {
			product *= primes[higherIndex]
		}
	}
	return product
}