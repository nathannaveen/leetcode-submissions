func maximumValue(strs []string) int {
    res := 0

    for _, str := range strs {
        n, err := strconv.Atoi(str)

        if err != nil {
            res = max(res, len(str)) 
        } else {
            res = max(res, n)
        }
    }

    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}