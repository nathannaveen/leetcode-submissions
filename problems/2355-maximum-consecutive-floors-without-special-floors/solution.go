func maxConsecutive(bottom int, top int, special []int) int {
    sort.Ints(special)
    
    special = append([]int{bottom - 1}, special...)
    special = append(special, top + 1)
    
    max := 0
    
    for i := 0; i < len(special) - 1; i++ {
        temp := special[i + 1] - special[i]
        if temp > max {
            max = temp - 1
        }
    }
    
    return max
}