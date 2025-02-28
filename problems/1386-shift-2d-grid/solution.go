func shiftGrid(grid [][]int, k int) [][]int {
    row := len(grid)
    col := len(grid[0])
    if k == row * col { return grid }
    tmp := []int{}
    for i := 0; i < row; i++ {
        for j := 0; j < col; j++ {
            tmp = append(tmp, grid[i][j])
        }
    }
    k = k % (row * col)
    tmp = append(tmp[len(tmp) - k:], tmp[:len(tmp) - k]... )
    res := make([][]int, row)
    index := 0
    for i := range res {
        res[i] = make([]int, col)
    }
    for i := 0; i < row; i++ {
        for j := 0; j < col; j++ {
            res[i][j] = tmp[index]
            index++
        }
    }
    return res
}