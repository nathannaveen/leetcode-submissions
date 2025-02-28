func isAcronym(words []string, s string) bool {
    if len(words) != len(s) {
        return false
    }
    for i := 0; i < len(s); i++ {
        if s[i] == words[i][0] {
            continue
        }
        return false
    }
    return true
}