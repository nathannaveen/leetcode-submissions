func pivotInteger(n int) int {
    for i := 1; i <= n; i++ {
        a := float64(i * (i + 1)) / float64(2)
        b := (float64(n - i + 1) / 2) * float64(i + n)

        if a == b {
            return i
        }
    }

    return -1
}