func canAliceWin(n int) bool {
    a := true
    for i := 10; i >= 0; i-- {
        n -= i
        if n < 0 {
            return !a
        }
        a = !a
    }

    return !a
}
