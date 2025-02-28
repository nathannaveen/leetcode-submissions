func doesAliceWin(s string) bool {
    for _, l := range s {
        if l == 'a' || l == 'e' || l == 'i' || l == 'o' || l == 'u' {
            return true
        }
    }

    return false
}
