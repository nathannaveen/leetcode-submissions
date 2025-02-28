func appendCharacters(s string, t string) int {
    j := 0

    for i := 0; i < len(s); i++ {
        if s[i] == t[j] {
            j++
        }
        if j == len(t) {
            break
        }
    }

    return len(t) - j
}