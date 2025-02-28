func partitionLabels(s string) []int {
    m := map[byte] []int {} // letter : []string{ first index, last index }
    res := []int{}
    prev := -1
    cnt := 0

    for i := 0; i < len(s); i++ {
        if _, ok := m[s[i]]; !ok {
            m[s[i]] = []int{i, i}
        }
        m[s[i]][1] = i
    }

    for i := 0; i < len(s); i++ {
        if m[s[i]][0] == i {
            cnt++
        }

        if m[s[i]][1] == i {
            cnt--
        }

        if cnt == 0 {
            res = append(res, i - prev)
            prev = i
        }
    }

    return res
}