func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	res := 0
	for i := 1; i < len(profit); i++ {
		if i >= 1 && profit[i-1] > profit[i] {
			profit[i-1], profit[i] = profit[i], profit[i-1]
			difficulty[i-1], difficulty[i] = difficulty[i], difficulty[i-1]
			i -= 2
		}
	}

	for _, ability := range worker {
		for i := len(profit) - 1; i >= 0; i-- {
			if difficulty[i] <= ability {
				res += profit[i]
				break
			}
		}
	}
	return res
}