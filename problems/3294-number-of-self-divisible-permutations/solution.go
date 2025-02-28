var dp = map[int] int {}

func selfDivisiblePermutationCount(n int) int {
    dp = map[int] int {}
    return helper(n, 0, 1)
}

func helper(n, used, cur int) int {
    if cur == n + 1 {
        return 1
    }

    if val, ok := dp[used]; ok {
        return val
    }

    res := 0

    for i := 0; i < n; i++ {
        if used & (1 << i) == 0 && gcd(i + 1, cur) == 1 {
            res += helper(n, used | (1 << i), cur + 1)
        }
    }

    dp[used] = res

    return res
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}