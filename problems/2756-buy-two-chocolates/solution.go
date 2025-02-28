func buyChoco(prices []int, money int) int {
    sort.Ints(prices)
    k := prices[0] + prices[1]
    if money - k >= 0 {
        return money - k
    }
    return money
}