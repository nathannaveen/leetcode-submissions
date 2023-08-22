func repeatedSubstringPattern(s string) bool {
    x := s + s

    return strings.Contains(x[1: len(x) - 1], s)
}