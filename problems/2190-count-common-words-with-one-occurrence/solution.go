func countWords(words1 []string, words2 []string) int {
    m := make(map[string] int)
    doubleInSameArray := make(map[string] bool)
    res := 0
    
    for _, word := range words1 {
        m[word]++
        if m[word] == 2 { doubleInSameArray[word] = true }
    }
    
    for _, word := range words2 { m[word]-- }
    
    for key, val := range m { if val == 0 && !doubleInSameArray[key] { res++ } }
    
    return res
}