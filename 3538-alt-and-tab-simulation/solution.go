func simulationResult(windows []int, queries []int) []int {
    res := make([]int, len(windows))
    cur := 0
    m := map[int] bool {}

    for i := len(queries) - 1; i >= 0; i-- {
        q := queries[i]
        if !m[q] {
            res[cur] = q
            cur++
            m[q] = true
        }
    }

    for _, window := range windows {
        if !m[window] {
            res[cur] = window
            cur++
        }
    }

    return res
}
