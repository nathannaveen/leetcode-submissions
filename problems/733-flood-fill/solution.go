type key struct {
    y int
    x int
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
    if image[sr][sc] != newColor {
        stack := []key{ key{sr, sc} }
        val := image[sr][sc]

        for len(stack) != 0 {
            pop := stack[len(stack) - 1]
            stack = stack[:len(stack) - 1]

            if pop.x > 0 && image[pop.y][pop.x - 1] == val {
                stack = append(stack, key{ pop.y, pop.x - 1 })
            }
            if pop.x < len(image[0]) - 1 && image[pop.y][pop.x + 1] == val {
                stack = append(stack, key{ pop.y, pop.x + 1 })
            }

            if pop.y > 0 && image[pop.y - 1][pop.x] == val {
                stack = append(stack, key{ pop.y - 1, pop.x })
            }
            if pop.y < len(image) - 1 && image[pop.y + 1][pop.x] == val {
                stack = append(stack, key{ pop.y + 1, pop.x })
            }

            image[pop.y][pop.x] = newColor
        }
    }
    
    return image
}