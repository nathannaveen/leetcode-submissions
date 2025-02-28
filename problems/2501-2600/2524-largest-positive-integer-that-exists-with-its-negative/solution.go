func findMaxK(nums []int) int {
    max := -1
    m := map[int] bool {}
    
    for _, num := range nums {
        if m[-num] && abs(num) > max {
            max = abs(num)
        }
        m[num] = true
    }
    
    return max
}

func abs(a int) int {
    if a > 0 {
        return a
    }
    return -a
}