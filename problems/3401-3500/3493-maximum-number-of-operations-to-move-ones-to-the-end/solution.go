func maxOperations(s string) int {
    res := 0
    oneCnt := 0

    for i := 0; i < len(s) - 1; i++ {
        if s[i] == '1' {
            oneCnt++
        }

        j := i + 1
        for j < len(s) && s[i] == '1' && s[j] == '0' {
            j++
        }
        j--
        
        if j != i {
            i = j
            res += oneCnt
        }
    }

    return res
}
