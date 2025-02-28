func numberCount(a int, b int) int {
    res := 0

    for i := a; i <= b; i++ {
        digits := map[int] bool {}
        res++

        n := i
        for n > 0 {
            x := n % 10
            if digits[x] {
                res--
                break
            }
            digits[x] = true
            n /= 10
        }
    }

    return res
}