func maximumTotalSum(maximumHeight []int) int64 {
    sort.Slice(maximumHeight, func(i, j int) bool {
        return maximumHeight[i] > maximumHeight[j]
    })

    res := int64(0)
    m := maximumHeight[0] + 1

    for _, h := range maximumHeight {
        v := min(m - 1, h)

        if v == 0 {
            return -1
        }

        res += int64(v)
        m = v
    }

    return res
}
