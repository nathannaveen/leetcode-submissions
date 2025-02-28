func minChanges(s string) int {
    res := 0

    for i := 0; i < len(s) - 1; i += 2 {
        if s[i] != s[i + 1] {
            res++
        }
    }

    return res
}