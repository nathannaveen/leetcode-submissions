func finalPrices(prices []int) []int {
	index := 0
	minimum := 1001
	
	for i, price := range prices {
		hasMinimum := true
		if index <= i {
			hasMinimum = false
			for j := i + 1; j < len(prices); j++ {
				if prices[j] <= price {
					hasMinimum = true
					minimum = prices[j]
					break
				}
			}
		}

		if hasMinimum {
			prices[i] = price - minimum
		}
	}
	return prices
}