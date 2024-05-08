func countSubstrings(s string, c byte) int64 {
    n := 0

    for _, l := range s {
        if c == byte(l) {
            n++
        }
    }

    return int64(n) * int64(n + 1) / 2
}
