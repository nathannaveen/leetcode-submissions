type p struct {
    i, j int
}

func queensAttacktheKing(queens [][]int, king []int) [][]int {
    directions := []p{ p{0, 1}, p{0, -1}, p{1, 0}, p{-1, 0}, p{1, 1}, p{-1, -1}, p{-1, 1}, p{1, -1} }
    
    m := map[p] bool {}
    
    for _, queen := range queens {
        m[p{ queen[0], queen[1] }] = true
    }
    
    res := [][]int{}
    
    for k := 0; k < 8; k++ {
        i := king[0]
        j := king[1]
        
        I := directions[k].i
        J := directions[k].j
        
        for i >= 0 && i < 8 && j >= 0 && j < 8 {
            i += I
            j += J
            if m[p{ i, j }] {
                res = append(res, []int{ i, j })
                break
            }
        }
    }
    
    return res
}