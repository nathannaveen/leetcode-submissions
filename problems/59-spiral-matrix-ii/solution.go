type point struct {
    i, j int
}

func generateMatrix(n int) [][]int {
    moves := []point{
        point{ 0, 1 },
        point{ 1, 0 },
        point{ 0, -1 },
        point{ -1, 0 },
    }
    
    res := make([][]int, n)
    for i := 0; i < n; i++ {
        res[i] = make([]int, n)
    }
    
    p := point{ 0, 0 }
    c := 0
    
    
    for k := 0; k < n * n - 1; k++ {
        move := moves[c % len(moves)]
        i := p.i + move.i
        j := p.j + move.j
        
        if i < 0 || j < 0 || i >= n || j >= n || res[i][j] != 0 {
            c++
            k--
        } else {
            res[p.i][p.j] = k + 1
            p = point{ i, j }
        }
    }
    
    res[p.i][p.j] = n * n
    
    return res
}