func maxCount(m int, n int, ops [][]int) int {
    x, y := m, n
    for _, op := range ops {
        x = min(op[0], x)
        y = min(op[1], y)
    }
    return x * y
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}