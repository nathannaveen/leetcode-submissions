func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	boats := 0
	left, right := 0, len(people)-1

	for left < right {
		if people[left]+people[right] <= limit {
			boats++
			left++
			right--
		} else {
			boats++
			right--
		}
	}
	if left == right {
		boats ++
	}
	return boats
}