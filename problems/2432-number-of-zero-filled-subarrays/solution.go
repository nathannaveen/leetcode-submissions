func zeroFilledSubarray(nums []int) int64 {
    n := int64(0)
    res := int64(0)
    
    for _, num := range nums {
        if num == 0 {
            n++
        } else {
            res += (n * (n + 1)) / 2
            n = 0
        }
    }
    
    res += (n * (n + 1)) / 2
    
    return res
}