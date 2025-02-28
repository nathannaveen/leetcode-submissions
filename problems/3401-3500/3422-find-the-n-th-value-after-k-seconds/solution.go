func valueAfterKSeconds(n int, k int) int {
    arr := make([][]int, k + 1)
    
    for i := 0; i < k + 1; i++ {
        x := make([]int, n)
        for j := 0; j < n; j++ {
            x[j] = 1
        }
        arr[i] = x
    }
    
    for i := 1; i <= k; i++ {
        for j := 1; j < n; j++ {
            arr[i][j] = (arr[i - 1][j] + arr[i][j - 1]) % 1000000007
        }
    }
    
    return arr[k][n - 1] % 1000000007
}
