var dp = map[int] int {}

type key struct {
    a, b int
}

func maxScore(nums []int) int {
    dp = map[int] int {}
    return helper(nums, 0, 1, map[key] int {})
}

func helper(nums []int, used int, cnt int, gcds map[key] int) int {
    res := 0

    if val, ok := dp[used]; ok {
        return val
    }
    
    for i := 0; i < len(nums); i++ {
        if used & (1 << i) != 0 {
            continue
        }

        for j := i + 1; j < len(nums); j++ {
            if used & (1 << j) != 0 {
                continue
            }

            if _, ok := gcds[key{nums[i], nums[j]}]; !ok {
                gcds[key{nums[i], nums[j]}] = GCD(nums[i], nums[j])
            }

            newMask := used | (1 << i) | (1 << j)
            gcd := gcds[key{nums[i], nums[j]}]
            
            res = max(res, cnt * gcd + helper(nums, newMask, cnt + 1, gcds))
        }
    }

    dp[used] = res

    return res
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}