func winnerSquareGame(n int) bool {
    dp := make([]bool, n + 1)
    
    for i := 1; i <= n; i++ {
        for j := 1; j * j <= i; j++ {
            if !dp[i - j * j] {
                dp[i] = true
                break
            }
        }
    }
    return dp[n]
}