func numberOfSubstrings(s string) int64 {
    m := make(map[rune] int)
    res := int64(0)
    
    for _, letter := range s {
        m[letter]++
    }
    
    for _, val := range m {
        res += int64((val * (val + 1)) / 2)
    }
    
    return res
}