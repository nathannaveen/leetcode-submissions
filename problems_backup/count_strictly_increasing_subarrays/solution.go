func countSubarrays(nums []int) int64 {
    res := 0
    
    prev := 0
    n := 0
    
    for _, num := range nums {
        if num > prev {
            prev = num
            n++
        } else {
            res += (n * (n + 1)) / 2
            prev = num
            n = 1
        }
    }
    
    res += (n * (n + 1)) / 2
    
    return int64(res)
}