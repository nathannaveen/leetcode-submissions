func minChanges(n int, k int) int {
    res := 0

    for n > 0 || k > 0 {
        if (n & 1) != (k & 1) {
            if (n & 1) == 0 {
                return -1
            } else {
                res++
            }
        }
        n >>= 1
        k >>= 1
    }

    return res
}
