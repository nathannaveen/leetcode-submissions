func minProductSum(nums1 []int, nums2 []int) int {
	res := 0
	sort.Ints(nums1)
	sort.Sort(sort.Reverse(sort.IntSlice(nums2)))

	for i := 0; i < len(nums1); i++ {
		res += nums1[i] * nums2[i]
	}
	return res
}