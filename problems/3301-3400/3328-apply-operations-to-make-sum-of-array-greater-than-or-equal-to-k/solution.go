func minOperations(k int) int {
    res := 100000

    for i := 1; i <= k; i++ {
        x := i + int(math.Ceil(float64(k) / float64(i))) - 2

        if x < res {
            res = x
        }
    }

    return res
}
