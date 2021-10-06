var m = make(map[byte] int)

func isAlienSorted(words []string, order string) bool {
    m = make(map[byte] int)
    for i := range order {
        m[order[i]] = i
    }
    for i := 1; i < len(words); i++ {
        word1, word2 := words[i - 1], words[i]
        
        // for j := 0; j < len(word1); j++ {
        //     if j >= len(word2) { return false }
        //     if m[word1[j]] > m[word2[j]] { 
        //         return false 
        //     }
        // }
        if !temp(word1, word2) { return false }
    }
    return true
}

func temp(word1, word2 string) bool {
    
    for i := 0; i < len(word1); i++ {
        if i >= len(word2) {
            return false
        }
        if m[word1[i]] != m[word2[i]] { 
            return m[word1[i]] <= m[word2[i]] 
        }
    }
    
    return true
}