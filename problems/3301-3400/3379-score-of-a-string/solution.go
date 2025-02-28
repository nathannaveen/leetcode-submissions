func scoreOfString(s string) int {
    res := 0

    for i := 0; i < len(s) - 1; i++ {
        res += abs(int(s[i]) - int(s[i + 1]))
    }

    return res
}

func abs(n int) int {
    if n > 0 {
        return n
    }
    return -n
}
