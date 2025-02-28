func maxScoreIndices(nums []int) []int {
    total := 0
    
    for _, n := range nums {
        total += n
    }
    
    max := total
    res := []int{0}
    
    for i, n := range nums {
        if n == 0 {
            total++
        } else {
            total--
        }
        
        if total > max {
            max = total
            res = []int{ i + 1 }
        } else if total == max {
            res = append(res, i + 1)
        }
        
    }
    return res
}