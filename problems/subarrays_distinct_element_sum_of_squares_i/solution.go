func sumCounts(nums []int) int {
    res := 0

    for i := 0; i < len(nums); i++ {
        distinct := make(map[int] bool)
        for j := i; j < len(nums); j++ {
            distinct[nums[j]] = true
            res += len(distinct) * len(distinct)
        }
    }

    return res
}