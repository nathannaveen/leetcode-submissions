func longestPalindrome(words []string) int {
    wordsSameLetters := map[string] int {}
    numMirroredWords := 0
    wordsMap := map[string] int {}
    
    for _, word := range words {
        s := string(word[1]) + string(word[0])
        
        if wordsMap[s] > 0 {
            numMirroredWords += 2
            wordsMap[s]--
            if wordsSameLetters[s] > 0 {
                wordsSameLetters[s]--
                if wordsSameLetters[s] == 0 {
                    delete(wordsSameLetters, s)
                }
            }
        } else if s == word {
            wordsSameLetters[word]++
            wordsMap[word]++
        } else {
            wordsMap[word]++
        }
    }
    
    res := numMirroredWords * 2
    
    if len(wordsSameLetters) >= 1 {
        res += 2
    }
    
    return res
}