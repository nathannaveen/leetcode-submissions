func countValidWords(sentence string) int {
    res := 0
    
    arr := strings.Fields(sentence)
    
    for _, word := range arr {
        canDo := true
        hyphens := 0
        punct := 0
        
        for i, l := range word {
            if unicode.IsDigit(rune(l)) {
                canDo = false
                break
            }
            if l == '-' {
                hyphens++
                if hyphens > 1 {
                    canDo = false
                    break
                }
                if i == 0 || (i != 0 && !unicode.IsLetter(rune(word[i - 1]))) ||
                i == len(word) - 1 || (i != len(word) - 1 && !unicode.IsLetter(rune(word[i + 1]))) {
                    canDo = false
                    break
                }
            } else if unicode.IsPunct(rune(l)) {
                punct++
                if punct > 1 || i != len(word) - 1 {
                    canDo = false
                    break
                }
            }
        }
        
        if canDo {
            res++
        }
    }
    
    return res
}