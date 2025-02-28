type pos struct {
    i, j int
}

func countBlackBlocks(m int, n int, coordinates [][]int) []int64 {
    res := make([]int64, 5)
    res[0] = int64((m - 1) * (n - 1))

    blocks := map[pos] int {} // the top left corner : the number of black cells

    for _, c := range coordinates {
        for _, p := range []pos{ {c[0], c[1]}, {c[0] - 1, c[1] - 1}, {c[0] - 1, c[1]}, {c[0], c[1] - 1} } {
            if p.i < 0 || p.j < 0 || p.i >= m - 1 || p.j >= n - 1 {
                continue
            }
            res[blocks[p]]--
            blocks[p]++
            res[blocks[p]]++
        }
    }

    return res
}