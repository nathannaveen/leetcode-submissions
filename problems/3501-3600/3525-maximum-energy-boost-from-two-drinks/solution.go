type key struct {
    index int
    x int
}

func maxEnergyBoost(energyDrinkA []int, energyDrinkB []int) int64 {
    dp := map[key] int64 {}
    return max(helper(energyDrinkA, energyDrinkB, 0, 0, dp), helper(energyDrinkA, energyDrinkB, 0, 1, dp))
}

func helper(a []int, b []int, index int, x int, dp map[key] int64) int64 {
    // x means which array we are in, 0 for a, 1 for b
    if index == len(a) {
        return 0
    }

    if val, ok := dp[key{index, x}]; ok {
        return val
    }

    res := int64(0)

    if x == 0 {
        res = max(res, int64(a[index]) + helper(a, b, index + 1, x, dp))
        res = max(res, helper(a, b, index + 1, 1, dp))
    } else {
        res = max(res, int64(b[index]) + helper(a, b, index + 1, x, dp))
        res = max(res, helper(a, b, index + 1, 0, dp))
    }

    dp[key{index, x}] = res

    return res
}

func max(a, b int64) int64 {
    if a > b {
        return a
    }
    return b
}
