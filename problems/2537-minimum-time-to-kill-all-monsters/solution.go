import "math"

var dp = map[int] int64 {}

func minimumTime(power []int) int64 {
    dp = map[int] int64 {}
    return helper(power, 0, 1, 0)
}

func helper(power []int, monstersDefeated, gain int, cnt int) int64 {
    if val, ok := dp[monstersDefeated]; ok {
        return val
    }

    res := int64(10000000000)

    for i := 0; i < len(power); i++ {
        if monstersDefeated & (1 << i) == 0 {
            x := int64(math.Ceil(float64(power[i]) / float64(gain))) + helper(power, monstersDefeated | (1 << i), gain + 1, cnt + 1)

            if x < res {
                res = x
            }
        }
    }

    if res == 10000000000 {
        res = 0
    }

    dp[monstersDefeated] = res

    return res
}