func generatePossibleNextMoves(currentState string) []string {
    res := []string{}
    
    for i := 1; i < len(currentState); i++ {
        cur, prev := currentState[i], currentState[i - 1]
        
        if cur == '+' && cur == prev {
            res = append(res, currentState[: i - 1] + "--" + currentState[i + 1:])
        }
    }
    
    return res
}