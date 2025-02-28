func findPrefixScore(nums []int) []int64 {
    max := int64(0)
    res := []int64{0}
    
    for _, n := range nums {
        if int64(n) > max {
            max = int64(n)
        }
        res = append(res, res[len(res) - 1] + max + int64(n))
    }
    
    return res[1:]
}