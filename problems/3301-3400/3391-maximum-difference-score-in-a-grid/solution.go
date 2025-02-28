type key struct {
    i, j int
}

var res = 0
var dp = map[key] int {}
var c = false

func maxScore(grid [][]int) int {
    res = 0
    dp = map[key] int {}
    c = false
    helper(grid, 0, 0, 0)
    return res
}

func helper(grid [][]int, i, j, cnt int) int {
    if i >= len(grid) || j >= len(grid[0]) {
        return 0
    }
    
    if val, ok := dp[key{i, j}]; ok {
        return val
    }
    
    x := max(helper(grid, i + 1, j, cnt + 1), helper(grid, i, j + 1, cnt + 1))
    z := x - grid[i][j]
    
    // fmt.Println(i, j, x, z, res)
    
    if z > res || (cnt == 1 && !c) {
        res = z
        c = true
    }
    
    dp[key{i, j}] = max(x, grid[i][j])
    
    return dp[key{i, j}]
}

/*
[6,5,1],
[5,7,9],
[6,7,4],
[6,10,5]
*/
