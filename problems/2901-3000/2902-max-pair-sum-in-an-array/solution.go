func maxSum(nums []int) int {
    res := -1

    for i := 0; i < len(nums); i++ {
        d := maxDigit(nums[i])
        for j := i + 1; j < len(nums); j++ {
            d2 := maxDigit(nums[j])

            if d == d2 {
                res = max(res, nums[i] + nums[j])
            }
        }
    }

    return res
}

func maxDigit(n int) int {
    d := 0

    for n > 0 {
        d = max(d, n % 10)
        n /= 10
    }

    return d
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}