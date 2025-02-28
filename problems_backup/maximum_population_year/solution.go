func maximumPopulation(logs [][]int) int {
	populationPerYear := make([]int, 101)
	res := 0
	for _, log := range logs {
		populationPerYear[log[0] - 1950]++
		populationPerYear[log[1] - 1950]--
	}
	for i := 1; i < 101; i++ {
		populationPerYear[i] += populationPerYear[i - 1]
		if populationPerYear[i] > populationPerYear[res] {
			res = i
		}
	}
	return res + 1950
}
