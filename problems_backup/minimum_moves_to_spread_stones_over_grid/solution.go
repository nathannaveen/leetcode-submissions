type key struct {
    i, j int
}

func minimumMoves(grid [][]int) int {
    extras := map[key] int {}
    empties := map[key] int {}
    
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid); j++ {
            if grid[i][j] > 1 {
                extras[key{i, j}] = grid[i][j] - 1
            } else if grid[i][j] == 0 {
                empties[key{i, j}] = 0
            }
        }
    }
    
    fmt.Println(extras)
    fmt.Println(empties)
    
    return helper(extras, empties)
}

func helper(extras, empties map[key] int) int {
    
    for extra, v := range extras {
        if v == 0 {
            continue
        }
        min := 1000
        for empty, _ := range empties {
            newExtras := map[key] int {}
            for k, v := range extras {
                newExtras[k] = v
            }
            newEmpties := map[key] int {}
            for k, v := range empties {
                newEmpties[k] = v
            }
            newExtras[extra]--
            delete(newEmpties, empty)

            d := abs(extra.j - empty.j) + abs(extra.i - empty.i)
            
            x := helper(newExtras, newEmpties) + d
            if x < min {
                min = x
            }
        }
        
        return min
    }
    
    return 0
}

func abs(a int) int {
    if a > 0 {
        return a
    }
    return -a
}


