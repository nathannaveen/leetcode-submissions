func findTheLongestBalancedSubstring(s string) int {
    search := ""
    max := 0

    for len(search) <= len(s) {
        if strings.Contains(s, search) {
            max = len(search)
        }
        search = "0" + search + "1"
    }

    return max
}