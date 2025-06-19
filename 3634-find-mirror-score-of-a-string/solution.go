func calculateScore(s string) int64 {
    res := int64(0)
    m := map[byte] []int {}

    for i := 0; i < len(s); i++ {
        x := 'z' - (s[i] - 'a')
        if _, ok := m[x]; ok {
            res += int64(i - m[x][len(m[x]) - 1])
            m[x] = m[x][:len(m[x]) - 1]
            if len(m[x]) == 0 {
                delete(m, x)
            }
        } else {
            m[s[i]] = append(m[s[i]], i)
        }
    }

    return res
}
