type key struct {
    i int
    val int
}

func canSeePersonsCount(heights []int) []int {
    res := make([]int, len(heights))
    stack := []key{}
    
    for i := 0; i < len(heights); i++ {
        for len(stack) > 0 && heights[i] > stack[len(stack) - 1].val {
            res[stack[len(stack) - 1].i]++
            stack = stack[:len(stack) - 1]
        }
        if len(stack) != 0 {
            res[stack[len(stack) - 1].i]++
        }
        stack = append(stack, key{ i, heights[i] })
    }
    
    return res
}