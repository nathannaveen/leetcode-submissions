type point struct {
    y int
    x int
}

func islandPerimeter(grid [][]int) int {
    res := 0
    stack := []point{}
    m := make(map[point] int)
    
    for i := 0; i < len(grid); i++ {
        equalsOne := false
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] == 1 {
                stack = append(stack, point{ i, j })
                equalsOne = true
                break
            }
        }
        if equalsOne {
            break
        }
    }
    
    for len(stack) != 0 {
        pos := stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
        m[pos]++
        
        if m[pos] == 1 {
            if pos.x > 0 && grid[pos.y][pos.x - 1] == 1 && m[point{ pos.y, pos.x - 1}] == 0 {
                stack = append(stack, point{ pos.y, pos.x - 1 })
            } else if m[point{ pos.y, pos.x - 1}] == 0 {
                res++
            }

            if pos.x < len(grid[0]) - 1 && grid[pos.y][pos.x + 1] == 1 && m[point{ pos.y, pos.x + 1}] == 0 {
                stack = append(stack, point{ pos.y, pos.x + 1 })
            } else if m[point{ pos.y, pos.x + 1}] == 0 {
                res++
            }

            if pos.y > 0 && grid[pos.y - 1][pos.x] == 1 && m[point{ pos.y - 1, pos.x}] == 0 {
                stack = append(stack, point{ pos.y - 1, pos.x })
            } else if m[point{ pos.y - 1, pos.x}] == 0 {
                res++
            }

            if pos.y < len(grid) - 1 && grid[pos.y + 1][pos.x] == 1 && m[point{ pos.y + 1, pos.x}] == 0 {
                stack = append(stack, point{ pos.y + 1, pos.x })
            } else if m[point{ pos.y + 1, pos.x}] == 0 {
                res++
            }
        }
    }
    
    return res
}