func triangularSum(nums []int) int {
    n := len(nums)
    for i := 0; i < len(nums) - 1; i++ {
        for j := 1; j < n; j++ {
            nums[j - 1] = (nums[j - 1] + nums[j]) % 10
        }
        n--
    }
    return nums[0]
}