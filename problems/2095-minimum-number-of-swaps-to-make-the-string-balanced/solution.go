func minSwaps(s string) int {
    stack := []string{}
    n := 0
    
    for i := 0; i < len(s); i++ {
        if s[i] == '[' {
            stack = append(stack, "[")
        } else {
            if len(stack) != 0 && stack[len(stack) - 1] == "[" {
                stack = stack[:len(stack) - 1] // pop of stack
            } else {
                stack = append(stack, "]")
                n++
            }
        }
    }
    
    return (n + 1) / 2
}