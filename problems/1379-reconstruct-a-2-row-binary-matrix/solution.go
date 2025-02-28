func reconstructMatrix(upper int, lower int, colsum []int) [][]int {
    res := [][]int{ make([]int, len(colsum)), make([]int, len(colsum)) }
    
    for i, c := range colsum {
        if c == 2 {
            res[0][i], res[1][i] = 1, 1
            upper--; lower--
        } else if c == 1 {
            if upper > lower {
                res[0][i] = 1
                upper--
            } else {
                res[1][i] = 1
                lower--
            }
        }
    }
    
    if upper != 0 || lower != 0 {
        return [][]int{}
    }
    
    return res
}