func arithmeticTriplets(nums []int, diff int) int {
    m := map[int] bool {}
    res := 0
    
    for _, n := range nums {
        if m[n - diff] && m[n - 2 * diff] {
            res++
        }
        m[n] = true
    }
    
    return res
}