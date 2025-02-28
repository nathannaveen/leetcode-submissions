func finalValueAfterOperations(operations []string) int {
    res := 0
    
    for _, operation := range operations {
        if operation[1] == '-' { 
            res-- 
            continue
        }
        res++
    }
    return res
}