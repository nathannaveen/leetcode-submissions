func winningPlayerCount(n int, pick [][]int) int {
    m := map[int] map[int] int {}
    res := 0
    done := map[int] bool {}

    for _, p := range pick {
        if _, ok := m[p[0]]; !ok {
            m[p[0]] = map[int] int {}
        }
        if done[p[0]] {
            continue
        }
        m[p[0]][p[1]]++
        if m[p[0]][p[1]] > p[0] {
            res++
            done[p[0]] = true
        }
    }

    return res
}
