func duplicateZeros(arr []int) {
	h := []int{}
	arrCounter := 0
	for _, i := range arr {
		h = append(h, i)
	}

	for _, i2 := range h {
		if arrCounter == len(h) {
			break
		}

		if i2 != 0 {
			arr[arrCounter] = i2
		} else {
			arr[arrCounter] = i2
			arrCounter++
			if arrCounter == len(h) {
				break
			}
			arr[arrCounter] = i2
		}
		arrCounter++
	}
}