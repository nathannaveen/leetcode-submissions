func areOccurrencesEqual(s string) bool {
    max := 0
    counter := 0
    m := make(map[string] int)
    
    for i := 0; i < len(s); i++ {
        m[string(s[i])]++
        
        if m[string(s[i])] == max {
            counter++
        } else if m[string(s[i])] > max {
            counter = 1
            max = m[string(s[i])]
        }
    }
    
    return len(m) == counter
}