func numberOfSpecialSubstrings(s string) int {
    lettersIndex := map[byte] int {} // letter : index
    cnt := 0
    res := 0

    for i := 0; i < len(s); i++ {
        val, contains := lettersIndex[s[i]]

        if contains && val >= i - cnt {
            cnt = i - lettersIndex[s[i]] - 1
        }
        lettersIndex[s[i]] = i
        cnt++
        res += cnt
    }

    return res
}
