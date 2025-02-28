func maximumNumberOfStringPairs(words []string) int {
    m := map[string] int {}
    res := 0

    for _, word := range words {
        revWord := reverse(word)
        if m[revWord] > 0 {
            res++
            m[revWord]--
            continue
        }
        m[word]++
    }

    return res
}

func reverse(s string) string {
    x := ""

    for _, l := range s {
        x = string(l) + x
    }

    return x
}