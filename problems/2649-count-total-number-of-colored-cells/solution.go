func coloredCells(n int) int64 {
    res := int64(1)

    if n == 1 {
        return 1
    }

    for i := 1; i <= n; i++ {
        res += int64(i + (i - 1) + ((i - 1) + (i - 2)))
    }

    return res
}
