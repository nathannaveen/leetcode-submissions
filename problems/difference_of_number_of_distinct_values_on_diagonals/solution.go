type pos struct {
    i, j int
}

type key struct {
    prevD int
    index int
}

func differenceOfDistinctValues(grid [][]int) [][]int {
    res := [][]int{}
    
    for i := 0; i < len(grid); i++ {
        res = append(res, make([]int, len(grid[0])))
        for j := 0; j < len(grid[0]); j++ {
            d := map[int] bool {}
            for k := 1; k <= min(i, j); k++ {
                d[grid[i - k][j - k]] = true
            }
            topLeft := len(d)
            d = map[int] bool {}
            for k := 1; k < min(len(grid) - i, len(grid[0]) - j); k++ {
                d[grid[i + k][j + k]] = true
            }
            bottomRight := len(d)
            // fmt.Println(i, j, topLeft, bottomRight)
            res[i][j] = abs(topLeft - bottomRight)
        }
    }
    
    return res
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}