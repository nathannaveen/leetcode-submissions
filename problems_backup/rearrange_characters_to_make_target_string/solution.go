func rearrangeCharacters(s string, target string) int {
    m := map[string] int {}
    
    m2 := map[string] int {}
    
    for _, letter := range s {
        m[string(letter)]++
    }
    
    for _, letter := range target {
        m2[string(letter)]++
    }
    
    min := 100000
    
    for a, b := range m2 {
        if m[a] / b < min {
            min = m[a] / b
        }
    }
    
    return min
}