func numberOfSubstrings(s string) int64 {
    m := make(map[string] int)
    res := int64(0)
    
    for _, letter := range s {
        m[string(letter)]++
        res += int64(m[string(letter)])
    }
    
    return res
}