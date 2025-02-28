func calPoints(ops []string) int {
    res := 0
    stack := []int{}
    
    for _, opp := range ops {
        if opp == "+" {
            res += stack[len(stack) - 1] + stack[len(stack) - 2]
            stack = append(stack, stack[len(stack) - 1] + stack[len(stack) - 2])
        } else if opp == "D" {
            res += stack[len(stack) - 1] * 2
            stack = append(stack, stack[len(stack) - 1] * 2)
        } else if opp == "C" {
            res -= stack[len(stack) - 1]
            stack = stack[:len(stack) - 1]
        } else {
            val, _ := strconv.Atoi(opp)
            res += val
            stack = append(stack, val)
        }
    }
    
    return res
}