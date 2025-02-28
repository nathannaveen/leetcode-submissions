func prefixCount(words []string, pref string) int {
    res := 0
    for _, word := range words {
        if len(word) >= len(pref) && word[:len(pref)] == pref {
            res++
        }
    }
    return res
}