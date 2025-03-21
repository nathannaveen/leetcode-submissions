func smallestNumber(n int, t int) int {
    for i := n; i <= n * (n / 10 + 1) * 10; i++ {
        prod := 1
        x := i
        for x > 0 {
            prod *= x % 10
            x /= 10
        }
        if prod % t == 0 {
            return i
        }
    }

    return -1
}
