func subarraySum(nums []int) int {
    res := 0
    sum := 0
    prefix := make([]int, len(nums) + 1)

    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        prefix[i + 1] = sum
        start := max(0, i - nums[i])

        res += prefix[i + 1] - prefix[start]
    }

    return res
}
