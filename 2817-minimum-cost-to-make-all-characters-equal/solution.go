func minimumCost(s string) int64 {
    res := int64(0)
    
    for i := 0; i < len(s) / 2 - 1; i++ {
        if s[i] != s[i + 1] {
            res += int64(i + 1)
        }
    }
    
    for i := len(s) - 1; i > len(s) / 2; i-- {
        if s[i] != s[i - 1] {
            res += int64(len(s) - i)
        }
    }
    if len(s) > 1 && s[len(s) / 2 - 1] != s[len(s) / 2] {
        res += int64(len(s) / 2)
    }
    
    return res
}