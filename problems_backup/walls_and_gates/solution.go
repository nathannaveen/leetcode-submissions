type key struct {
    r int // row
    c int // col
}

func wallsAndGates(rooms [][]int)  {
    queue := []key{}
    maxRow := len(rooms)
    maxCol := len(rooms[0])
    alreadyDid := make(map[key] bool)
    d := 0
    
    isProperRoom := func (r, c int, rooms [][]int) {
    if r >= 0 && c >= 0 && r < maxRow && c < maxCol && alreadyDid[key{ r, c }] == false && rooms[r][c] != -1 {
        alreadyDid[key{ r, c }] = true
        queue = append(queue, key{ r, c })
    }
}
    
    for i := 0; i < maxRow; i++ {
        for j := 0; j < maxCol; j++ {
            if rooms[i][j] == 0 {
                queue = append(queue, key{ i, j })
                alreadyDid[key{ i, j }] = true
            }
        }
    }
    
    for len(queue) != 0 {
        n := len(queue)
        
        for i := 0; i < n; i++ {
            pop := queue[0]
            queue = queue[1:]
            row := pop.r
            col := pop.c
            rooms[row][col] = d
            
            isProperRoom(row + 1, col, rooms)
            isProperRoom(row - 1, col, rooms)
            isProperRoom(row, col + 1, rooms)
            isProperRoom(row, col - 1, rooms)
        }
        d++
    }
}

