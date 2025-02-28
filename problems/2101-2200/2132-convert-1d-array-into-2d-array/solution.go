func construct2DArray(original []int, m int, n int) [][]int {
    res := [][]int{}
    
    if len(original) == m * n {
        for i := 0; i < m; i++ {
            res = append(res, original[i * n : (i + 1) * n])
        }
    }
    
    return res
}