// Define the directions to search in (up, down, left, right)
var directions = []pos{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

type pos struct {
    i, j int
}

func findMaxFish(grid [][]int) int {
    // Keep track of visited cells to avoid processing them again
    visited := make(map[pos] bool)
    max := 0

    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            x := dfs(grid, pos{i, j}, visited)

            if x > max {
                max = x
            }
        }
    }

    return max
}

// Depth-first search function to search for fish starting from a given cell
func dfs(grid [][]int, cur pos, visited map[pos]bool) int {
    // Check if the current cell is out of bounds or 
    // has been visited before or does not contain any fish
    if !inBounds(grid, cur) || grid[cur.i][cur.j] == 0 || visited[cur] {
        return 0
    }

    // Add the number of fish in the current cell to the result
    res := grid[cur.i][cur.j]

    // Mark the current cell as visited
    visited[cur] = true

    // Recursively search in all adjacent cells
    for _, dir := range directions {
        res += dfs(grid, pos{cur.i + dir.i, cur.j + dir.j}, visited)
    }

    // Return the total number of fish caught in this search
    return res
}

// Function to check if a given cell is within the bounds of the grid
func inBounds(grid [][]int, cur pos) bool {
    if cur.i < 0 || cur.j < 0 || cur.i >= len(grid) || cur.j >= len(grid[0]) {
        return false
    }
    
    return true
}
