func suggestedProducts(products []string, searchWord string) [][]string {
	res := make([][]string, len(searchWord))
	sort.Strings(products)
	s := ""

	for i, i2 := range searchWord {
		s += string(i2)
		counter := 0
		for _, product := range products {
			if counter == 3 {
				break
			}
			if strings.HasPrefix(product, s) {
				counter++
				res[i] = append(res[i], product)
			}
		}

	}

	return res
}
