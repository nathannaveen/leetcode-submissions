func checkAlmostEquivalent(word1 string, word2 string) bool {
    m := make(map[int] int)
    
    for i := 0; i < len(word1); i++ {
        m[int(word1[i] - 'a')]++
        m[int(word2[i] - 'a')]--
    }
    
    for _, val := range m {
        if val < -3 || val > 3 { return false }
    }
    
    return true
}