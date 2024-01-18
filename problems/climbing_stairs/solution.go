func climbStairs(n int) int {
    dp := map[int] int {}
    return helper(n, 0, dp)
}

func helper(n, cur int, dp map[int] int) int {
    if val, ok := dp[cur]; ok {
        return val
    }

    if n == cur {
        return 1
    }

    res := 0

    if n - cur >= 1 {
        res += helper(n, cur + 1, dp)
    }
    if n - cur >= 2 {
        res += helper(n, cur + 2, dp)
    }

    dp[cur] = res

    return res
}