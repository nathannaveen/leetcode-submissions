func minOperations(nums []int, queries []int) []int64 {
	n := len(nums)
	sort.Ints(nums)

	prefixSums := make([]int, n)
	sum := 0
	
	for i := 0; i < n; i++ {
		sum += nums[i]
		prefixSums[i] = sum
	}

	res := make([]int64, len(queries))

	for i, query := range queries {
		idx := sort.SearchInts(nums, query)
		if idx == 0 {
			res[i] = int64(prefixSums[n-1] - query*n)
		} else {
			res[i] = int64(prefixSums[n-1] - query*(n-idx) + query*idx - prefixSums[idx-1]<<1)
		}
	}
	return res
}