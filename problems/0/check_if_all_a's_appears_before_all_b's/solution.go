func checkString(s string) bool {
    splitS := strings.Split(s, "")
    sort.Strings(splitS)
    return s == strings.Join(splitS, "")
}