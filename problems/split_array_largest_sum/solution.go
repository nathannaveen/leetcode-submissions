type key struct {
    a, b int
}

var dp = map[key] int {}

func splitArray(nums []int, m int) int {
    dp = map[key] int {}

    prefix := make([]int, len(nums) + 1)
    prefix[0] = 0

    for i := 0; i < len(nums); i++ {
        prefix[i + 1] = prefix[i] + nums[i]
    }

    return helper(prefix, m, 1)
}

func helper(prefix []int, m, cur int) int {
    // return minimum

    if m == 1 {
        return prefix[len(prefix) - 1] - prefix[cur - 1]
    }

    if val, ok := dp[key{cur, m}]; ok {
        return val
    }

    min := math.MaxInt

    for i := cur; i < len(prefix) - 1; i++ {
        x := max(helper(prefix, m - 1, i + 1), prefix[i] - prefix[cur - 1])

        if x < min {
            min = x
        }
    }

    dp[key{cur, m}] = min

    return min
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}