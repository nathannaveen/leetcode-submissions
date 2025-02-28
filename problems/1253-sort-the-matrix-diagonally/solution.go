func diagonalSort(mat [][]int) [][]int {
    for i := 0; i < len(mat)-1; i++ {
        for j := 0; j < len(mat[0])-1; j++ {
            newI := i + 1
            newJ := j + 1
            for k := 1; k <= min(newI, newJ); k++ {
                if mat[newI-k][newJ-k] > mat[newI-k+1][newJ-k+1] {

                    mat[newI-k][newJ-k], mat[newI-k+1][newJ-k+1] = 
                    mat[newI-k+1][newJ-k+1], mat[newI-k][newJ-k]
                }
            }
        }
    }
    return mat
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}