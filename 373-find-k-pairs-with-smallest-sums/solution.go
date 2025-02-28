func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	type a struct {
		key1, key2, sum int
	}
	arr := []a{}
	res := [][]int{}

	for _, i := range nums1 {
		for _, i2 := range nums2 {
			arr = append(arr, a{
				i, i2, i + i2,
			})
		}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].sum < arr[j].sum
	})

	for i := 0; i < int(math.Min(float64(k), float64(len(arr)))); i++ {
		res = append(res, []int{arr[i].key1, arr[i].key2})
	}
	return res
}