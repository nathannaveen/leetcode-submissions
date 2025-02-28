func findLUSlength(a string, b string) int {
    if a == b {
        return -1
    }
    return max(a, b)
}

func max(a, b string) int {
    if len(a) > len(b) {
        return len(a)
    }
    return len(b)
}