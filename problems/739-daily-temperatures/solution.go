type key struct {
    val int
    index int
}

func dailyTemperatures(T []int) []int {
    stack := []key{}
    res := make([]int, len(T))
    
    for i := len(T) - 1; i >= 0; i-- {
        
        counter := 0
        
        for j := len(stack) - 1; j >= 0 && T[i] >= stack[j].val; j-- {
            counter++ // Add to counter so we don't keep on re-making the stack
        }
        
        stack = stack[:len(stack) - counter]
        res[i] = 0
        
        if len(stack) != 0 {
            res[i] = stack[len(stack) - 1].index - i
        }
        
        stack = append(stack, key{ T[i], i })
    }
    
    return res
}