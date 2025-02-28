func maxProfit(prices []int) int {
	dp := make([]int, len(prices))
    last := prices[len(prices) - 1]
    x := 0

    for i := len(prices) - 1; i >= 0; i-- {
        if prices[i] > last {
            last = prices[i]
        }

        if last - prices[i] > x {
            x = last - prices[i]
        }

        dp[i] = x
    }

    x = prices[0]

    max := 0

    for i := 0; i < len(prices); i++ {
        if prices[i] > x {
            if i != len(prices) - 1 {
                if prices[i] - x + dp[i + 1] > max {
                    max = prices[i] - x + dp[i + 1]
                }
            }
            if prices[i] - x > max {
                max = prices[i] - x
            }
        } else {
            x = prices[i]
        }
    }

    return max
}
