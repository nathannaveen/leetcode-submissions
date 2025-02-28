func multiply(mat1 [][]int, mat2 [][]int) [][]int {
    res := make([][]int, len(mat1))
    
    for i := 0; i < len(mat1); i++ {
        res[i] = make([]int, len(mat2[0]))
        
        for j := 0; j < len(mat2[0]); j++ {
            for k := 0; k < len(mat1[0]); k++ {
                res[i][j] += mat1[i][k] * mat2[k][j]
            }
        }
    }
    
    return res
}