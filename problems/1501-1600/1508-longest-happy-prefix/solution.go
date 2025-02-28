func longestPrefix(s string) string {
    i := 0
    max := ""
    for i < len(s) - 1 {
        if s[0 : i + 1] == s[len(s) - i - 1 : len(s)] {
            max = s[0 : i + 1]
        }
        
        i++
    }
    return max
}