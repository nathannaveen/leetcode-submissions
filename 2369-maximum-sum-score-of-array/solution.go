func maximumSumScore(nums []int) int64 {
    lSum, rSum := int64(0), int64(0)
    res := int64(-100000)
    
    for i := 0; i < len(nums); i++ {
        lSum += int64(nums[i])
        rSum += int64(nums[len(nums) - 1 - i])
        
        res = max(res, max(lSum, rSum))
    }
    return res
}

func max(a, b int64) int64 {
    if a > b { return a }
    return b
}