func rangeAddQueries(n int, queries [][]int) [][]int {
    windows := make([][]int, n)

    for i := range windows {
        windows[i] = make([]int, n + 1)
    }

    for _, query := range queries {
        for row := query[0]; row <= query[2]; row++ {
            windows[row][query[1]]++
            windows[row][query[3] + 1]--
        }
    }

    res := make([][]int, n)

    for i := 0; i < len(res); i++ {
        sum := 0
        res[i] = make([]int, n)

        for j := 0; j < n; j++ {
            sum += windows[i][j]
            res[i][j] = sum
        }
    }

    return res
}