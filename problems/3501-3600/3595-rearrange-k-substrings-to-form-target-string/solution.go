func isPossibleToRearrange(s string, t string, k int) bool {
    m := map[string] int {}
    div := len(s) / k

    for i := 0; i < len(s); i += div {
        m[s[i : i + div]]++
    }

    for i := 0; i < len(t); i += div {
        x := t[i : i + div]
        m[x]--
        if m[x] == 0 {
            delete(m, x)
        }
    }

    return len(m) == 0
}
