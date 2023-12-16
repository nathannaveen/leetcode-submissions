func isAnagram(s string, t string) bool {
    m := map[byte] int {}

    for i := 0; i < len(t); i++ {
        m[t[i]]++
    }

    for i := 0; i < len(s); i++ {
        m[s[i]]--

        if m[s[i]] == 0 {
            delete(m, s[i])
        }
    }

    return len(m) == 0
}