func getSmallestString(s string) string {
    for i := 0; i < len(s) - 1; i++ {
        n, _ := strconv.Atoi(string(s[i]))
        n2, _ := strconv.Atoi(string(s[i + 1]))

        if n % 2 == n2 % 2 && n2 < n {
            return s[:i] + string(s[i + 1]) + string(s[i]) + s[i + 2:]
        }
    }

    return s
}
