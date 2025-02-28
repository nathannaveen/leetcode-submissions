func minOperations(s string) int {
    res := 0
    
    for i := 0; i < len(s); i++ {
        if int(s[i]) - 48 != i % 2 {
            res++
        }
    }
    
    return int(math.Min(float64(res), float64(len(s) - res)))
}