func rangeBitwiseAnd(left int, right int) int {
    res := left

    for i := left + 1; i <= right; i++ {
        res &= i
        if res == 0 {
            return 0
        }
    }

    return res
}