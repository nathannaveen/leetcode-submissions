func modifiedMatrix(matrix [][]int) [][]int {
    for j := 0; j < len(matrix[0]); j++ {
        neg := []int{}
        m := 0

        for i := 0; i < len(matrix); i++ {
            m = max(m, matrix[i][j])
            if matrix[i][j] == -1 {
                neg = append(neg, i)
            }
        }

        for _, n := range neg {
            matrix[n][j] = m
        }
    }

    return matrix
}