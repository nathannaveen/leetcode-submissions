func largestSubarray(nums []int, k int) []int {
	maximumSubarray := make([]int, k)
	for i := 0; i < len(nums) - k + 1; i++ {
		currentSubarray := nums[i : i + k]
		maximumSubarray = whichArrayIsGreater(maximumSubarray, currentSubarray)
	}
	return maximumSubarray
}

func whichArrayIsGreater(one, two []int) []int {
	for i := 0; i < len(one); i++ {
		if one[i] > two[i] {
			return one
		} else if one[i] < two[i] {
			return two
		}
	}
	return []int{}
}