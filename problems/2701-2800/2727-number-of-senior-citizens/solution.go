func countSeniors(details []string) int {
    res := 0

    for _, d := range details {
        n, _ := strconv.Atoi(d[11:13])

        if n > 60 {
            res++
        }
    }

    return res
}