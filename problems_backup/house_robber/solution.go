var dp map[int] int // index in nums : maximum value obtained

func rob(nums []int) int {
    dp = map[int] int {}

    return helper(-2, nums)
}

func helper(cur int, nums []int) int {
    max := 0

    if val, ok := dp[cur]; ok {
        return val
    }

    for i := cur + 2; i < len(nums); i++ {
        x := helper(i, nums)
        if x > max {
            max = x
        }
    }

    res := max

    if cur < len(nums) && cur >= 0 {
        res += nums[cur]
    }
    dp[cur] = res

    return res
}