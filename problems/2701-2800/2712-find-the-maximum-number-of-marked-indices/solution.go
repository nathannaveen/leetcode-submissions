func maxNumOfMarkedIndices(nums []int) int {
    sort.Ints(nums)

    i, j := 0, len(nums) / 2
    res := 0

    for i < len(nums) / 2 && j < len(nums) {
        if nums[i] * 2 <= nums[j] {
            i++
            res += 2
        }
        j++
    }

    return res
}