func greatestLetter(s string) string {
    m := make([]bool, 'z'+1)
    res := ""
    
    for _, letter := range s {
        m[letter] = true
    }
    
    for i := 90; i >= 65; i-- {
        if m[i] && m[i+32] {
            res = string(i)
            break
        }
    }
    
    return res
}