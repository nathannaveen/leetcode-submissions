func isGood(nums []int) bool {
    if len(nums) < 2 {
        return false
    }

    sort.Ints(nums)

    n := len(nums) - 1

    for i := 0; i < len(nums) - 2; i++ {
        if nums[i] != i + 1 {
            return false
        }
    }

    return nums[len(nums) - 2] == n && nums[len(nums) - 1] == n
}