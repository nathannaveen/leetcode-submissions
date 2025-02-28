var res int

func minIncrements(n int, cost []int) int {
    res = 0
    dfs(1, n, cost, 0)
    return res
}

func dfs(cur, n int, cost []int, depth int) int {
    // return the cost
    if cur * 2 > n && cur * 2 + 1 > n {
        return cost[cur - 1]
    }
    
    left := 0
    right := 0
    
    if cur * 2 <= n {
        left = dfs(cur * 2, n, cost, depth + 1)
    }
    
    if cur * 2 + 1 <= n {
        right = dfs(cur * 2 + 1, n, cost, depth + 1)
    }
    
    res += abs(left - right)
    
    return cost[cur - 1] + max(left, right)
}

func abs(a int) int {
    if a > 0 {
        return a
    }
    return -a
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}