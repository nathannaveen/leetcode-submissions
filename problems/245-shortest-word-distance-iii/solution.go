func shortestWordDistance(wordsDict []string, word1 string, word2 string) int {
    minimum := len(wordsDict) - 1
    prev := -1
    prevWord := ""
    
    for i, word := range wordsDict {
        if (word == word1 && prevWord == word2) || (word == word2 && prevWord == word1) {
            
            if i - prev < minimum { 
                minimum = i - prev 
            }
            
            prevWord = word
            prev = i
        } else if word == word1 || word == word2 { 
            prev = i 
            prevWord = word
        }
    }
    return minimum
}