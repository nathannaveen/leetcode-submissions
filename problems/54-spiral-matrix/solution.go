type point struct {
    i, j int
}

func spiralOrder(matrix [][]int) []int {
    n, m := len(matrix), len(matrix[0])
    
    moves := []point{ 
        point{ 0, 1 },
        point{ 1, 0 },
        point{ 0, -1 },
        point{ -1, 0 },
    }
    
    res := make([]int, n * m)
    
    p := point{ 0, 0 }
    c := 0
    
    for k := 0; k < n * m - 1; k++ {
        move := moves[c % len(moves)]
        i := p.i + move.i
        j := p.j + move.j
        
        if i < 0 || j < 0 || i >= n || j >= m || matrix[i][j] == -101 {
            c++
            k--
        } else {
            res[k] = matrix[p.i][p.j]
            matrix[p.i][p.j] = -101
            p = point{ i, j }
        }
    }
    
    res[len(res) - 1] = matrix[p.i][p.j]
    
    return res
}