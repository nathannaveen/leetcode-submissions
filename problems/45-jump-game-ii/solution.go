var dp = map[int] int {}

func jump(nums []int) int {
    dp = map[int] int {}

    return helper(nums, 0)
}

func helper(nums []int, i int) int {
    // return the min

    if i >= len(nums) - 1 {
        return 0
    }

    if val, ok := dp[i]; ok {
        return val
    } else {
        dp[i] = 10001
    }

    for j := 1; j <= nums[i]; j++ {
        dp[i] = min(dp[i], 1 + helper(nums, i + j))
    }

    return dp[i]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}