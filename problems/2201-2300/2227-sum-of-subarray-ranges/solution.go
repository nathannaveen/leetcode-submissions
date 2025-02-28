func subArrayRanges(nums []int) int64 {
    res := int64(0)
    for i := 0; i < len(nums); i++ {
        min, max := nums[i], nums[i]
        for j := i; j < len(nums); j++ {
            if nums[j] < min {
                min = nums[j]
            }
            if nums[j] > max {
                max = nums[j]
            }
            res += int64(max - min)
        }
    }
    return res
}