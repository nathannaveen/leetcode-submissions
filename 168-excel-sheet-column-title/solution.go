func convertToTitle(columnNumber int) string {
    res := ""

    for columnNumber > 0 {
        columnNumber--
        res = string(rune(columnNumber % 26 + 'A')) + res
        columnNumber /= 26
    }

    return res
}