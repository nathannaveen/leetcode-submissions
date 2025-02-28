func findPairs(nums []int, k int) int {
    m := make(map[int] int)
    res := 0
    
    for _, n := range nums {
        m[n]++
    }
    
    for n := range m {
        if (k > 0 && m[n + k] > 0) || (k == 0 && m[n] > 1) {
            res++
        }
    }
    
    return res
}