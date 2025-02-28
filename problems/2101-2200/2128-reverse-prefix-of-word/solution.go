func reversePrefix(word string, ch byte) string {
    s := ""
    
    for i, letter := range word {
        s = string(letter) + s
        
        if letter == rune(ch) {
            return s + word[i + 1 :]
        }
    }
    
    return word
}