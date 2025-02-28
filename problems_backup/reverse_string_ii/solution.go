func reverseStr(s string, k int) string {
    rev := true
    res := ""
    for i := 0; i < len(s); i += k {
        if rev {
            res += reverse(s[i : min(i + k, len(s))])
        } else {
            res += s[i : min(i + k, len(s))]
        }

        rev = !rev
    }

    return res
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func reverse(s string) string {
    s2 := ""

    for _, l := range s {
        s2 = string(l) + s2
    }

    return s2
}