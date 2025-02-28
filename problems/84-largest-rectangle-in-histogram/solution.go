type key struct {
    index int // start width of the rectangle
    val int // height of the rectangle
}

func largestRectangleArea(heights []int) int {
    newHeights := make([]int, len(heights))
    copy(newHeights, heights)
    newHeights = append(newHeights, -1)
    
    stack := []key{}
    maxArea := 0
    
    for i, height := range newHeights {
        index := i
        for len(stack) != 0 && stack[len(stack) - 1].val > height {
            pop := stack[len(stack) - 1]
            stack = stack[:len(stack) - 1]
            index = pop.index
            if pop.val * (i - index) > maxArea {
                maxArea = pop.val * (i - index) 
            }
        }
        stack = append(stack, key{ index, height })
    }
    
    return maxArea
}