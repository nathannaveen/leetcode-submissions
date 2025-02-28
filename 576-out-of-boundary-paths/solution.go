type key struct {
    i, j, maxMove int
}

var dp = map[key] int {}

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
    dp = map[key] int {}
    return helper(m, n, startRow, startColumn, maxMove)
}

func helper(m, n, i, j, maxMove int) int {
    if i == -1 || j == -1 || i == m || j == n {
        return 1
    }
    
    if maxMove == 0 {
        return 0
    }
    
    if val, ok := dp[key{ i, j, maxMove }]; ok {
        return val
    }
    
    res := 0
    
    res = (res + helper(m, n, i + 1, j, maxMove - 1)) % 1000000007
    res = (res + helper(m, n, i - 1, j, maxMove - 1)) % 1000000007
    res = (res + helper(m, n, i, j + 1, maxMove - 1)) % 1000000007
    res = (res + helper(m, n, i, j - 1, maxMove - 1)) % 1000000007
    
    dp[key{ i, j, maxMove }] = res
    
    return res
}