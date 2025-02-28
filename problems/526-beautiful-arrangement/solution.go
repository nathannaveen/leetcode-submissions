var dp = map[int] int {}

func countArrangement(n int) int {
    dp = map[int] int {}
    return helper(n, 0, 1)
}

func helper(n int, used int, cnt int) int {
    if cnt == n + 1 {
        return 1
    }

    if val, ok := dp[used]; ok {
        return val
    }

    res := 0

    for i := 1; i <= n; i++ {
        if used & (1 << i) == 0 && (i % cnt == 0 || cnt % i == 0) {
            res += helper(n, used | (1 << i), cnt + 1)
        }
    }
    
    dp[used] = res

    return res
}