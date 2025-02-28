func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] > people[j][1]
		}
		return people[i][0] < people[j][0]
	})

	res := make([][]int, len(people))
	empties := make([]int, len(people))

	for i := 0; i < len(people); i++ {
		empties[i] = i
	}

	for _, person := range people {
		res[empties[person[1]]] = person
		empties = append(empties[:person[1]], empties[person[1]+1:]...)
	}

	return res
}