func maxJump(stones []int) int {
    res := stones[1] - stones[0]

    for i := 0; i < len(stones) - 2; i++ {
        res = max(res, stones[i + 2] - stones[i])
    }

    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}