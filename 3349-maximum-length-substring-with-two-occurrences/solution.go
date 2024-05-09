func maximumLengthSubstring(s string) int {
    m := map[byte] int {}
    l := 0
    res := 0

    for i := 0; i < len(s); i++ {
        m[s[i]]++

        for m[s[i]] > 2 {
            m[s[i - l]]--
            l--
        }

        l++

        if l > res {
            res = l
        }
    }

    return res
}
