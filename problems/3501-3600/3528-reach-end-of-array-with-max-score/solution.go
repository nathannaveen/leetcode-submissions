func findMaximumScore(nums []int) int64 {
    max := int64(0)
    res := int64(0)
    for _, n := range nums {
        res += max
        if int64(n) > max {
            max = int64(n)
        }
    }

    return res
}
