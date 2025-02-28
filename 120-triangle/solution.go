func minimumTotal(triangle [][]int) int {
    n := len(triangle)
    
    for i := 1; i < n; i++ {
        m := len(triangle[i])
        
        for j := 0; j < m; j++ {
            
            sumOne := 10001
            sumTwo := 10001
            
            if j != 0 {
                sumOne = triangle[i - 1][j - 1] + triangle[i][j]
            }
            
            if j != m - 1 {
                sumTwo = triangle[i - 1][j] + triangle[i][j]
            }
            
            triangle[i][j] = int(math.Min(float64(sumOne), float64(sumTwo)))
            
        }
    }
    
    min := 2000001
    
    for j := 0; j < len(triangle[n - 1]); j++ {
        if triangle[n - 1][j] < min {
            min = triangle[n - 1][j]
        }
    }
    return min
}

