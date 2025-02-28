var dp = map[int] int {}

func maximumANDSum(nums []int, numSlots int) int {
    dp = map[int] int {}
    return helper(nums, 0, 0, numSlots)
}

func helper(nums []int, cur int, used int, numSlots int) int {
    if cur == len(nums) {
        return 0
    }

    if val, ok := dp[used]; ok {
        return val
    }

    res := 0

    for i := 0; i < numSlots; i++ {
        newUsed := used
        canDo := false

        if newUsed & (1 << i) == 0 {
            newUsed |= 1 << i
            canDo = true
        } else if newUsed & (1 << (i + numSlots)) == 0 {
            newUsed |= 1 << (i + numSlots)
            canDo = true
        }

        if canDo {
            res = max(res, nums[cur] & (i + 1) + helper(nums, cur + 1, newUsed, numSlots))
        }
    }

    dp[used] = res

    return res
}