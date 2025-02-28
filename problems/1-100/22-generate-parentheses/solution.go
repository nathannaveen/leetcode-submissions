func generateParenthesis(n int) []string {
    return helper("", n, n, n, []string{})
}

func helper(s string, numOpen, numClose, n int, res []string) []string {
    if numOpen == 0 && numClose == 0 {
        res = append(res, s)
        return res
    }
    
    if numOpen > 0 && numOpen - 1 < numClose {
        res = helper(s + "(", numOpen - 1, numClose, n, res)
    }
    if numClose > 0 {
        res = helper(s + ")", numOpen, numClose - 1, n, res)
    }
    return res
}