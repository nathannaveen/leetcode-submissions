func isPrefixString(s string, words []string) bool {
    temp := ""
    for _, i := range words {
        temp += i
        if temp == s {
            return true
        }
    }
    return false
}