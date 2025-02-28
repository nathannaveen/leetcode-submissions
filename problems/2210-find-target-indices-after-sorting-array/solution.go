func targetIndices(nums []int, target int) []int {
    startAt := 0
    c := 0
    res := []int{}
    
    for _, num := range nums {
        if num == target {
            c++
        } else if num < target {
            startAt++
        }
    }
    
    for i := 0; i < c; i++ {
        res = append(res, i + startAt)
    }
    
    return res
}