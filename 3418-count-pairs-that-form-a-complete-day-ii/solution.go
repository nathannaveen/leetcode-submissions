func countCompleteDayPairs(hours []int) int64 {
    diff := map[int] int64 {}
    res := int64(0)

    for _, hour := range hours {
        x := hour % 24

        if x == 0 {
            res += diff[0]
        } else {
            res += diff[24 - x]
        }

        diff[x]++
    }

    return res
}
