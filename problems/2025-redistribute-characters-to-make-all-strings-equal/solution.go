func makeEqual(words []string) bool {
    m := map[string] int {}
    
    for _, word := range words {
        for _, letter := range word {
            m[string(letter)]++
        }
    }
    
    for _, b := range m {
        if b % len(words) != 0 { return false }
    }
    
    return true
}