func checkStrings(s1 string, s2 string) bool {
    even := map[rune] int {}
    odd := map[rune] int {}

    for i, l := range s1 {
        if i % 2 == 0 {
            even[l]++
        } else {
            odd[l]++
        }
    }

    for i, l := range s2 {
        if i % 2 == 0 {
            even[l]--
            if even[l] == 0 {
                delete(even, l)
            }
        } else {
            odd[l]--
            if odd[l] == 0 {
                delete(odd, l)
            }
        }
    }

    return len(even) == 0 && len(odd) == 0
}