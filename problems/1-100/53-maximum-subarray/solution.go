var dp = map[int] int {}

func maxSubArray(nums []int) int {
    dp = map[int] int {}
    res := -10000

    for i := 0; i < len(nums); i++ {
        x := maxSum(nums, i)
        if x > res {
            res = x
        }
    }

    return res
}

func maxSum(nums []int, i int) int {
    if i == len(nums) {
        return -10000
    }

    if val, ok := dp[i]; ok {
        return val
    }

    res := nums[i]

    x := maxSum(nums, i + 1)

    if x > 0 {
        res += x
    }

    dp[i] = res

    return res
}