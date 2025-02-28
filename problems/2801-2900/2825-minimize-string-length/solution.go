func minimizedStringLength(s string) int {
    m := map[string] int {}
    
    for i := 0; i < len(s); i++ {
        m[string(s[i])]++
    }
    return len(m)
}