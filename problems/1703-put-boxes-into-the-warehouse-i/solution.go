func maxBoxesInWarehouse(boxes []int, warehouse []int) int {
	boxCounter := 0
	sort.Ints(boxes)
	minimums := []int{warehouse[0]}

	for i := 1; i < len(warehouse); i++ {
		if minimums[i-1] > warehouse[i] {
			minimums = append(minimums, warehouse[i])
		} else {
			minimums = append(minimums, minimums[i-1])
		}
	}

	for i := len(minimums) - 1; i >= 0; i-- {
		if boxCounter == len(boxes) {
			break
		}
		if minimums[i] >= boxes[boxCounter] {
			boxCounter++
		}
	}

	return boxCounter 
}