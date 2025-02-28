func canBeTypedWords(text string, brokenLetters string) int {
    m := make(map[rune] int)
    res := 0
    
    for _, i := range brokenLetters {
        m[i]++
    }
    
    hasBrokenLetters := false
    
    for _, letter := range text {
        if letter == ' ' {
            if !hasBrokenLetters {
                res++
            }
            hasBrokenLetters = false
            
        } else if m[letter] >= 1 {
            hasBrokenLetters = true
        }
    }
    
    if !hasBrokenLetters {
        res++
    }
    
    return res
}