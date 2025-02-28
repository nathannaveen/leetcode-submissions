type key struct {
    a, b int
}

var dp = map[key] int {}

func specialPerm(nums []int) int {
    dp = map[key] int {}
    return helper(0, 0, nums, 1)
}

func helper(used int, cnt int, nums []int, prev int) int {
    if cnt == len(nums) {
        return 1
    }

    if val, ok := dp[key{used, prev}]; ok {
        return val
    }

    res := 0

    for i := 0; i < len(nums); i++ {
        if (used & (1 << i)) == 0 && (prev % nums[i] == 0 || nums[i] % prev == 0) {
            res += helper(used | (1 << i), cnt + 1, nums, nums[i])
        }
    }
    res = res % 1000000007

    dp[key{used, prev}] = res

    return res
}