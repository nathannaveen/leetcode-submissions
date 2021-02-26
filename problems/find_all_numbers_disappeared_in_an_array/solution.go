func findDisappearedNumbers(nums []int) []int {

	m := make(map[int]int)
	arr := []int{}

	for _, v := range nums {
		m[v]++
	}

	for i := 0; i <= len(nums); i++ {
		if _, ok := m[i]; !ok {
			arr = append(arr, i)
		}
	}
	return arr[1:]
}
