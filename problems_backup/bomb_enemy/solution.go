type pos struct {
    i, j int
}

func maxKilledEnemies(grid [][]byte) int {
    enemiesInRow := map[pos] int {}
    enemiesInCol := map[pos] int {}
    placesForBombs := []pos{}
    
    for i := 0; i < len(grid); i++ {
        start := 0
        enemies := 0
        
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] == 'E' {
                enemies++
            } else if grid[i][j] == 'W' {
                for k := start; k < j; k++ {
                    if grid[i][k] == '0' {
                        enemiesInRow[pos{ i, k }] = enemies
                    }
                }
                enemies = 0
                start = j + 1
            } else {
                placesForBombs = append(placesForBombs, pos{ i, j })
            }
        }
        
        for k := start; k < len(grid[0]); k++ {
            if grid[i][k] == '0' {
                enemiesInRow[pos{ i, k }] = enemies
            }
        }
    }
    
    for j := 0; j < len(grid[0]); j++ {
        start := 0
        enemies := 0
        
        for i := 0; i < len(grid); i++ {
            if grid[i][j] == 'E' {
                enemies++
            } else if grid[i][j] == 'W' {
                for k := start; k < i; k++ {
                    if grid[k][j] == '0' {
                        enemiesInCol[pos{ k, j }] = enemies
                    }
                }
                enemies = 0
                start = i + 1
            }
        }
        
        for k := start; k < len(grid); k++ {
            if grid[k][j] == '0' {
                enemiesInCol[pos{ k, j }] = enemies
            }
        }
    }
    
    max := 0
    
    for _, p := range placesForBombs {
        temp := enemiesInRow[p] + enemiesInCol[p]
        if temp > max {
            max = temp
        }
    }
    
    return max
}