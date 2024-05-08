func isSubstringPresent(s string) bool {
    m := map[string] bool {}

    for i := 0; i < len(s) - 1; i++ {
        m[string(s[i]) + string(s[i + 1])] = true
        if m[string(s[i + 1]) + string(s[i])] {
            return true
        }
    }

    return false
}
