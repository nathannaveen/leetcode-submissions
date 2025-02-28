func groupAnagrams(strs []string) [][]string {
    m := map[string] int {} // sorted string : index in res
    res := [][]string{}

    for _, s := range strs {
        x := strings.Split(s, "")
        sort.Strings(x)
        z := strings.Join(x, "")

        if _, ok := m[z]; !ok {
            m[z] = len(res)
            res = append(res, []string{})
        }
        res[m[z]] = append(res[m[z]], s)
    }

    return res
}
