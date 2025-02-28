func rearrangeArray(nums []int) []int {
    pos, neg := []int{}, []int{}
    res := []int{}
    
    for _, num := range nums {
        if num > 0 {
            pos = append(pos, num)
            continue
        }
        neg = append(neg, num)
    }
    
    for i := 0; i < len(pos); i++ {
        res = append(res, pos[i], neg[i])
    }
    
    return res
}