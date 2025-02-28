func lastNonEmptyString(s string) string {
    freq := map[string] int {}
    max := 0
    res := ""

    for _, l := range s {
        freq[string(l)]++
        
        if freq[string(l)] > max {
            res = string(l)
            max = freq[string(l)]
        } else if freq[string(l)] == max {
            res += string(l)
        }
    }

    return res
}