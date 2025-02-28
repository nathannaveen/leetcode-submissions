func maxSpending(values [][]int) int64 {
    arr := []int{}
    res := int64(0)

    for i := 0; i < len(values); i++ {
        for j := 0; j < len(values[0]); j++ {
            arr = append(arr, values[i][j])
        }
    }

    sort.Ints(arr)

    for i, r := range arr {
        res += int64(i + 1) * int64(r)
    }

    return res
}