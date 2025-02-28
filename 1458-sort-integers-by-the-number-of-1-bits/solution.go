func sortByBits(arr []int) []int {
	h := [501][]int{}

	for _, i := range arr {
		n := i
		numBits := 0
		for n > 0 {
			numBits += n & 1
			n >>= 1
		}
		h[numBits] = append(h[numBits], i)
	}

	arrayCounter := 0

	for i := 0; i < len(h); i++ {
		sort.Ints(h[i])
		for j := 0; j < len(h[i]); j++ {
			arr[arrayCounter] = h[i][j]
			arrayCounter++
		}
		if arrayCounter == len(arr) {
			break
		}
	}

	return arr
}