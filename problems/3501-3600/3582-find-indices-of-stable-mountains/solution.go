func stableMountains(height []int, threshold int) []int {
    res := []int{}
    for i := 1; i < len(height); i++ {
        if height[i - 1] > threshold {
            res = append(res, i)
        }
    }

    return res
}
