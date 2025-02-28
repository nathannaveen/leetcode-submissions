func maximumStrongPairXor(nums []int) int {
    sort.Ints(nums)

    max := 0

    for i := 0; i < len(nums); i++ {
        for j := i; j < len(nums); j++ {
            x := nums[i] ^ nums[j]
            if nums[j] - nums[i] <= nums[i] && x > max {
                max = x
            }
        }
    }

    return max
}