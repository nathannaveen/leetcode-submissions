func findLonely(nums []int) []int {
    res := []int{}
    m := make(map[int] int)
    
    for _, num := range nums {
        m[num]++
    }
    
    for a, b := range m {
        _, isPlusOne := m[a + 1]
        _, isMinusOne := m[a - 1]
        
        if b == 1 && !isPlusOne && !isMinusOne {
            res = append(res, a)
        }
    }
    
    return res
}