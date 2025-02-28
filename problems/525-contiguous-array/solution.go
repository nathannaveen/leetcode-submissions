func findMaxLength(nums []int) int {
    res := 0
    n := 0
    visited := make(map[int] int) // value : index
    visited[0] = -1
    
    for i, num := range nums {
        if num == 0 {
            n--
        } else {
            n++
        }
        
        if _, ok := visited[n]; ok {
            res = int(math.Max(float64(res), float64(i - visited[n])))
        } else {
            visited[n] = i
        }
    }
    
    return res
}