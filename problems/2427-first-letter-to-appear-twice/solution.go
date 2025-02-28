func repeatedCharacter(s string) byte {
    found := map[byte] bool {}
    
    for i := 0; i < len(s); i++ {
        letter := s[i]
        if found[letter] {
            return letter
        }
        found[letter] = true
    }
    return 'z'
}