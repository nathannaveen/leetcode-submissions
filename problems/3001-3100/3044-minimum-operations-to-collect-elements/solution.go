func minOperations(nums []int, k int) int {
    m := map[int] bool {}

    for i := len(nums) - 1; i >= 0; i-- {
        if nums[i] <= k {
            m[nums[i]] = true
        }

        if len(m) == k {
            return len(nums) - i
        }
    }

    return 0
}