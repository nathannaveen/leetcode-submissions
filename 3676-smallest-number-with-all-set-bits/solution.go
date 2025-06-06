func smallestNumber(n int) int {
    res := 1
    for res < n {
        res = res * 2 + 1
    }
    return res
}
