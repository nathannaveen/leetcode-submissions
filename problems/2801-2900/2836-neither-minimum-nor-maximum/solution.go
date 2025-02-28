func findNonMinOrMax(nums []int) int {
    sort.Ints(nums)

    if len(nums) > 2 {
        return nums[1]
    }
    return -1
}