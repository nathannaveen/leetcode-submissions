func findMinimumOperations(s1 string, s2 string, s3 string) int {
    m := map[string] int {}

    helper(m, s1)
    helper(m, s2)
    helper(m, s3)

    res := -1

    fmt.Println(m)

    for a, b := range m {
        if b == 3 {

            if len(a) > res {
                res = len(a)
            }
        }
    }

    if res == -1 {
        return -1
    }

    return len(s1) - res + len(s2) - res + len(s3) - res
}

func helper(m map[string] int, s1 string) {
    s := ""
    for _, l := range s1 {
        s += string(l)
        m[s]++
    }
}