type pos struct {
    i, j int
}

func firstCompleteIndex(arr []int, mat [][]int) int {
    m := map[int] pos {}
    
    for i := 0; i < len(mat); i++ {
        for j := 0; j < len(mat[0]); j++ {
            m[mat[i][j]] = pos{i, j}
        }
    }
    
    rows := make([]int, len(mat))
    cols := make([]int, len(mat[0]))
    
    for i := 0; i < len(arr); i++ {
        rows[m[arr[i]].i]++
        if rows[m[arr[i]].i] == len(mat[0]) {
            return i
        }
        
        cols[m[arr[i]].j]++
        if cols[m[arr[i]].j] == len(mat) {
            return i
        }
    }
    
    return 0
}