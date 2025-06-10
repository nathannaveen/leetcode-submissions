func constructTransformedArray(nums []int) []int {
    res := make([]int, len(nums))

    for i := 0; i < len(nums); i++ {
        res[i] = nums[(i + nums[i] % len(nums) + len(nums)) % len(nums)]
    }

    return res
}
