func wordPattern(pattern string, s string) bool {
    m := make(map[string] string) // letter : word
    m2 := make(map[string] string) // word : letter
    
    s2 := strings.Split(s, " ")
    if len(s2) != len(pattern) { 
        return false 
    }
    
    for i, word := range s2 {
        letter := string(pattern[i])
        
        if m[letter] == "" && m2[word] == "" { 
            // if not done this set of letter, word
            m[letter] = word
            m2[word] = letter
        } else if m[letter] != word || m2[word] != letter {
            // If we have already found another set of letter, word with a different combination
            return false
        }
    }
    
    return true
}