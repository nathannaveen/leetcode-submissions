func canPermutePalindrome(s string) bool {
    cntOdd := 0
    m := map[rune] int {}

    for _, l := range s {
        m[l]++
        if m[l] % 2 == 1 {
            cntOdd++
        } else {
            cntOdd--
        }
    }

    return cntOdd<= 1
}
