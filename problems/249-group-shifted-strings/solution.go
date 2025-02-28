func groupStrings(strings []string) [][]string {
    m := map[string] int {} // val : index in res
    res := [][]string{}

    for _, s := range strings {
        val := ""

        for i := 1; i < len(s); i++ {
            x := (26 + s[i] - s[i - 1]) % 26
            val += strconv.Itoa(int(x)) + " "
        }
        
        if _, ok := m[val]; !ok {
            res = append(res, []string{})
            m[val] = len(res) - 1
        }
        res[m[val]] = append(res[m[val]], s)
    }

    return res
}