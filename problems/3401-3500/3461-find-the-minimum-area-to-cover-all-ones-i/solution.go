func minimumArea(grid [][]int) int {
    maxI, minI, maxJ, minJ := 0, 1001, 0, 1001

    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] == 0 {
                continue
            }

            if i < minI {
                minI = i
            }
            if i > maxI {
                maxI = i
            }

            if j < minJ {
                minJ = j
            }
            if j > maxJ {
                maxJ = j
            }
        }
    }

    return (maxI - minI + 1) * (maxJ - minJ + 1)
}
