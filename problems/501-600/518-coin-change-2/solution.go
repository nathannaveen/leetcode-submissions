type key struct {
    a int
    i int
}

var c = []int{}
var dp = map[key] int {}

func change(amount int, coins []int) int {
    sort.Ints(coins)
    c = coins
    dp = map[key] int{}
    return helper(amount, 0)
}

func helper(amount, i int) int {
    if amount == 0 {
        return 1
    }
    if val, ok := dp[key{ amount, i }]; ok {
        return val
    }
    
    res := 0
    newI := i
    
    for newI < len(c) && c[newI] <= amount {
        res += helper(amount - c[newI], newI)
        newI++
    }
    
    dp[key{ amount, i }] = res
    
    return res
}