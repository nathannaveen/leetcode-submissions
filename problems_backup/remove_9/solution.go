func newInteger(n int) int {
    b := 1
    res := 0

    for n > 0 {
        res += n % 9 * b
        n /= 9
        b *= 10
    }

    return res
}
