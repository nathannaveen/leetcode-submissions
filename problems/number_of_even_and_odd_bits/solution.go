func evenOddBit(n int) []int {
    res := []int{0, 0}
    i := 0

    for n > 0 {
        if (n & 1) == 1 {
            if i % 2 == 0 { // even
                res[0]++
            } else { // odd
                res[1]++
            }
        }
        n >>= 1
        i++
    }

    return res
}