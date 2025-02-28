func findTheWinner(n int, k int) int {
    res := 0
    
    for i := 1; i <= n; i++ {
        res = (res + k) % i
    }
    
    return res + 1
}