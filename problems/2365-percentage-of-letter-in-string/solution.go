func percentageLetter(s string, letter byte) int {
    n := 0
    
    for i := 0; i < len(s); i++ {
        if s[i] == letter {
            n++
        }
    }
    
    return int(float64(n) / float64(len(s)) * 100)
}