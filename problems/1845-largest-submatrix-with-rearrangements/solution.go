func largestSubmatrix(matrix [][]int) int {
    n, m := len(matrix), len(matrix[0])
    heights := make([]int, m)
    res := 0
    
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if matrix[i][j] == 1 {
                heights[j]++
            } else {
                heights[j] = 0
            }
        }
        
        newHeights := make([]int, m)
        copy(newHeights, heights)
        sort.Ints(newHeights)
        
        for j := 0; j < m; j++ {
            temp := newHeights[j] * (m - j)
            if temp > res {
                res = temp
            }
        }
    }
    
    return res
}