func zigzagTraversal(grid [][]int) []int {
    res := []int{}
    forward := true
    on := true

    for i := 0; i < len(grid); i++ {
        j := 0
        x := 1
        end := len(grid[0])
        if !forward {
            j = len(grid[0]) - 1
            x = -1
            end = -1
        }

        for j != end {
            if on {
                res = append(res, grid[i][j])
            }
            on = !on
            j += x
        }

        forward = !forward
    }

    return res
}
