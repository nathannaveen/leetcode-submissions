func searchMatrix(matrix [][]int, target int) bool {
    i := 0
    
    for i = 0; i < len(matrix) - 1; i++ {
        if i + 1 < len(matrix) && matrix[i + 1][0] > target {
            break
        }
    }
    
    for j := 0; j < len(matrix[0]); j++ {
        if matrix[i][j] == target {
            return true
        } else if matrix[i][j] > target {
            break
        }
    }
    
    return false
}