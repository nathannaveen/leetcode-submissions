var dp = map[int] int {}

func maxScore(nums []int) int {
    dp = map[int] int {}

    return helper(nums, 0)
}

func helper(nums []int, cur int) int {
    max := 0

    if val, ok := dp[cur]; ok {
        return val
    }

    for i := cur + 1; i < len(nums); i++ {
        x := helper(nums, i) + (i - cur) * nums[i]
        if x > max {
            max = x
        }
    }

    dp[cur] = max

    return max
}
