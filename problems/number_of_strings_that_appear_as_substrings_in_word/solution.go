func numOfStrings(patterns []string, word string) int {
    res := 0
    for _, pat := range patterns {
        if strings.Index(word, pat) != -1 {
            res++
        }
    }
    return res
}