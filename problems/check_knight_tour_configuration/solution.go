type pos struct {
    i, j int
}

func checkValidGrid(grid [][]int) bool {
    arr := make([]pos, len(grid) * len(grid))

    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid); j++ {
            arr[grid[i][j]] = pos{i, j}
        }
    }

    if arr[0].i != 0 || arr[0].j != 0 {
        return false
    }

    for i := 1; i < len(arr); i++ {
        if !validMove(arr[i].i, arr[i].j, arr[i - 1].i, arr[i - 1].j) {
            return false
        }
    }

    return true
}

func validMove(si, sj, ei, ej int) bool {
    return (abs(si - ei) == 2 && abs(sj - ej) == 1) || (abs(si - ei) == 1 && abs(sj - ej) == 2)
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}