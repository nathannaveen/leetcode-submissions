type key struct {
    n int
    yStart int
    xStart int
}

func findFarmland(land [][]int) [][]int {
    arr := [][]key{}
    res := [][]int{}
    
    for i:= range land {
        arr = append(arr, make([]key, len(land[i])))
        
        for j := range land[i] {
            if land[i][j] == 0 {
                if i != 0 && land[i - 1][j] == 1 && (j == len(land[0]) - 1 || land[i - 1][j + 1] == 0) {
                    res = append(res, []int{ arr[i - 1][j].yStart, arr[i - 1][j].xStart, i - 1, j })
                }
                arr[i][j] = key{ 0, 0, 0 }
            } else {
                if (i == 0 || land[i - 1][j] == 0) && (j == 0 || land[i][j - 1] == 0) {
                    arr[i][j] = key{ 1, i, j }
                } else if i != 0 && land[i - 1][j] == 1 {
                    arr[i][j] = arr[i - 1][j]
                } else if j != 0 && land[i][j - 1] == 1 {
                    arr[i][j] = arr[i][j - 1]
                }
            }
        }
    }
    
    for i := range land[len(land) - 1] {
        if land[len(land) - 1][i] == 1 && (i == len(land[0]) - 1 || land[len(land) - 1][i + 1] == 0) {
            res = append(res, []int{ arr[len(land) - 1][i].yStart, arr[len(land) - 1][i].xStart, len(land) - 1, i })
        }
    }
    
    return res
}