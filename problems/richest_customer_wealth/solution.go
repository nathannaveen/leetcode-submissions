func maximumWealth(accounts [][]int) int {
    res := 0
    for _, account := range accounts {
        wealth := 0
        for _, money := range account {
            wealth += money
        }
        if wealth > res { res = wealth }
    }
    return res
}