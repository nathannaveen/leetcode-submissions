func decodeCiphertext(encodedText string, rows int) string {
    n := len(encodedText)
    cols := n / rows
    res := []string{}
    
    for i := 0; i < cols; i++ {
        for j := i; j < n; j += cols + 1 {
            res = append(res, string(encodedText[j]))
        }
    }
    
    return strings.TrimRight(strings.Join(res, ""), " ")
}