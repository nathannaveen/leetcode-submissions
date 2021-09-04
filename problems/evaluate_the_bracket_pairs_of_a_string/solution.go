func evaluate(s string, knowledge [][]string) string {
    storedKnowledge := make(map[string] string)
    
    for _, know := range knowledge {
        storedKnowledge[know[0]] = know[1]
    }
    
    arr := strings.Split(s, "(")
    res := arr[0]
    
    for i := 1; i < len(arr); i++ {
        temp := strings.Split(arr[i], ")")
        if val, ok := storedKnowledge[temp[0]]; ok {
            res += val
        } else {
            res += "?"
        }
        res += temp[1]
    }
    return res
}