func canAliceWin(nums []int) bool {
    s, d := 0, 0
    for _, n := range nums {
        if n > 9 {
            d += n
        } else {
            s += n
        }
    }

    return d != s
}
