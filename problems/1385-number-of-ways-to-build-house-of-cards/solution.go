type key struct {
    a, b int
}

var dp = map[key] int{}

func houseOfCards(n int) int {
    dp = map[key] int{}
    return helper(n, 500)
}

func helper(n, b int) int {
    if n == 0 || n == 2 {
        return 1
    }
    
    if val, ok := dp[key{ n, b }]; ok {
        return val
    }
    
    res := 0
    
    numTri := int(math.Min(float64((n - 2) / 3), float64(b - 1)))
    
    for i := 1; i <= numTri; i++ {
        res += helper(n - (i * 3) - 2, i)
    }
    
    dp[key{ n, b }] = res
    
    return dp[key{ n, b }]
}