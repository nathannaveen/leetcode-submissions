type key struct {
    x int
    y int
}

func minArea(image [][]byte, x int, y int) int {
    stack := []key{ key{ y, x } }
    minX, maxX, minY, maxY := 101, 0, 101, 0
    
    for len(stack) != 0 {
        pop := stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
        
        minX = min(pop.x, minX)
        maxX = max(pop.x, maxX)
        
        minY = min(pop.y, minY)
        maxY = max(pop.y, maxY)
        
        if pop.x > 0 && image[pop.y][pop.x - 1] == '1' {
            stack = append(stack, key{ pop.x - 1, pop.y })
            image[pop.y][pop.x - 1] = '0'
        }
        
        if pop.x < len(image[0]) - 1 && image[pop.y][pop.x + 1] == '1' {
            stack = append(stack, key{ pop.x + 1, pop.y })
            image[pop.y][pop.x + 1] = '0'
        }
        
        if pop.y > 0 && image[pop.y - 1][pop.x] == '1' {
            stack = append(stack, key{ pop.x, pop.y - 1 })
            image[pop.y - 1][pop.x] = '0'
        }
        
        if pop.y < len(image) - 1 && image[pop.y + 1][pop.x] == '1' {
            stack = append(stack, key{ pop.x, pop.y + 1 })
            image[pop.y + 1][pop.x] = '0'
        }
    }
    
    return (maxX - minX + 1) * (maxY - minY + 1)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a > b {
        return b
    }
    return a
}