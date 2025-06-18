func hasMatch(s string, p string) bool {
    arr := strings.Split(p, "*")

    x := strings.Index(s, arr[0])
    if x == -1 {
        return false
    }
    x += len(arr[0])
    return strings.Contains(s[x:], arr[1])
}
