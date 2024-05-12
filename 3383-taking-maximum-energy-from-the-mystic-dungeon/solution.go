func maximumEnergy(energy []int, k int) int {
    res := -1000
    dp := map[int] int {}
    
    for i := len(energy) - 1; i >= 0; i-- {
        x := dp[i + k] + energy[i]
        dp[i] = x
        if x > res {
            res = x
        }
    }
    
    return res
}
