func numKLenSubstrNoRepeats(s string, k int) int {
    if k > len(s) {
        return 0
    }

    res := 0

    numRepeats := 0
    m := map[byte] int {}

    for i := 0; i < len(s); i++ {
        if i >= k { // we can remove from the end
            m[s[i - k]]--
            if m[s[i - k]] == 1 {
                numRepeats--
            }
        }
        m[s[i]]++
        if m[s[i]] == 2 {
            numRepeats++
        }

        if numRepeats == 0 && i >= k - 1 {
            res++
        }
    }

    return res
}
