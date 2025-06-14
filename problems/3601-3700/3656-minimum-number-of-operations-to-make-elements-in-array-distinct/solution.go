func minimumOperations(nums []int) int {
    l := -1
    m := map[int] int {} // val : latestIndex

    for i, n := range nums {
        val, ok := m[n]
        if ok && val > l {
            l = val
        }
        m[n] = i
    }

    if l == -1 {
        return 0
    }

    return (l + (3 - l % 3)) / 3
}
