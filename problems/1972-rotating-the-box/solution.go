func rotateTheBox(box [][]byte) [][]byte {
    n, m := len(box), len(box[0])
    
    for row := 0; row < n; row++ {
        numStones := 0
        start := 0
        for i := 0; i < m; i++ {
            if box[row][i] == '#' {
                numStones++
            } else if box[row][i] == '*' {
                for j := i - numStones; j < i; j++ {
                    box[row][j] = '#'
                }
                for j := start; j < i - numStones; j++ {
                    box[row][j] = '.'
                }
                
                numStones = 0
                start = i + 1
            }
        }
        
        for j := m - numStones; j < m; j++ {
            box[row][j] = '#'
        }
        for j := start; j < m - numStones; j++ {
            box[row][j] = '.'
        }
    }
    
    res := make([][]byte, m)
    
    for i := 0; i < m; i++ {
        arr := make([]byte, n)
        for row := n - 1; row >= 0; row-- {
            arr[n - 1 - row] = box[row][i]
        }
        res[i] = arr
    }
    
    return res
}