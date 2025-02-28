func integerBreak(n int) int {
    res := 1

    if n == 2 {
        return 1
    } else if n == 3 {
        return 2
    }

    for n > 4 {
        res *= 3
        n -= 3
    }

    res *= n

    return res
}