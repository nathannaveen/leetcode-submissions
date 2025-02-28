func divideString(s string, k int, fill byte) []string {
    res := []string{}
    
    for len(s) > 0 {
        if len(s) >= k {
            res = append(res, s[:k])
            s = s[k:]
        } else {
            res = append(res, s)
            s = ""
        }
    }
    
    n := k - len(res[len(res) - 1])
    
    for i := 0; i < n; i++ {
        res[len(res) - 1] += string(fill)
    }
    return res
}