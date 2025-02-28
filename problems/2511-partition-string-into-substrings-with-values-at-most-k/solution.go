func minimumPartition(s string, k int) int {
    n, res := 0, 0

    for _, runeDigit := range s {
        digit := int(runeDigit) - 48

        n = n * 10 + digit

        if n > k {
            res++
            n = digit
            
            if n > k {
                return -1
            }
        }
    }

    return res + 1
}