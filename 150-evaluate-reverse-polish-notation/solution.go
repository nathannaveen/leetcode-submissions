func evalRPN(tokens []string) int {
    stack := []int{}
    
    for _, t := range tokens {
        if n, err := strconv.Atoi(t); err == nil {
            stack = append(stack, n)
        } else {
            a := stack[len(stack) - 2]
            b := stack[len(stack) - 1]
            stack = stack[:len(stack) - 2]
            
            switch t {
            case "+":
                stack = append(stack, a + b)
            case "-":
                stack = append(stack, a - b)
            case "*":
                stack = append(stack, a * b)
            case "/":
                stack = append(stack, a / b)
            }
        }
    }
    
    return stack[0]
}