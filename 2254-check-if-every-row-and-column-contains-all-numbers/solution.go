func checkValid(matrix [][]int) bool {
    n := len(matrix)
    rows := make([]map[int] bool, n)
    cols := make([]map[int] bool, n)
    counter := 0
    
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            val := matrix[i][j]
            
            if len(rows[i]) == 0 { rows[i] = make(map[int] bool) }
            rows[i][val] = true
            if len(rows[i]) == n { counter++ }
            
            if len(cols[j]) == 0 { cols[j] = make(map[int] bool) }
            cols[j][val] = true
            if len(cols[j]) == n { counter++ }
        }
    }
    
    return counter == 2 * n
}