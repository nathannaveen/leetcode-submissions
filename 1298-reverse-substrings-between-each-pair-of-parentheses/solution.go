func reverseParentheses(s string) string {
    stack := []string{}
    res := ""
    
    for _, i2 := range s {
        if i2 == '(' {
            stack = append(stack, "" )
        } else if i2 == ')' {
            pop := stack[len(stack) - 1]
            stack = stack[:len(stack) - 1]
            
            if len(stack) == 0 {
                res += reverse(pop)
            } else {
                stack[len(stack) - 1] += reverse(pop)
            }
        } else {
            if len(stack) == 0 {
                res += string(i2)
            } else {
                stack[len(stack) - 1] += string(i2)
            }
        }
    }
    
    return res
}

func reverse(s string) string {
    res := ""
    for _, i := range s {
        res = string(i) + res
    }
    return res
}