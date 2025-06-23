func minCost(arr []int, brr []int, k int64) int64 {
    res1 := int64(0)
    res2 := int64(0)
    for i := 0; i < len(arr); i++ {
        res1 += abs(arr[i] - brr[i])
    }
    sort.Ints(arr)
    sort.Ints(brr)
    for i := 0; i < len(arr); i++ {
        res2 += abs(arr[i] - brr[i])
    }
    return min(res1, res2 + k)
}

func abs(a int) int64 {
    if a < 0 {
        return int64(-a)
    }
    return int64(a)
}
