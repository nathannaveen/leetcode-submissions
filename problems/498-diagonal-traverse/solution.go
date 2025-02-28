func findDiagonalOrder(mat [][]int) []int {
    
    arr := make([][]int, len(mat[0]) * len(mat) - 1)
    if len(mat[0]) == 1 || len(mat) == 1 {
        arr = make([][]int, len(mat[0]) * len(mat))
    }
        
    for i := 0; i < len(mat); i++ {
        for j := 0; j < len(mat[0]); j++ {
            if (i + j) % 2 == 1 {
                arr[i + j] = append(arr[i + j], mat[i][j])
            } else {
                arr[i + j] = append(arr[i + j][:0], append([]int{mat[i][j]}, arr[i + j][0:]...)...)
            }
        }
    }
    res := []int{}
    
    for _, i := range arr {
        res = append(res, i...)
    }
    
    return res
}