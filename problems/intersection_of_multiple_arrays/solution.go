func intersection(nums [][]int) []int {
    m := map[int] int {}
    res := []int{}
    
    for _, arr := range nums {
        for _, num := range arr {
            m[num]++
        }
    }
    
    for a, b := range m {
        if b == len(nums) {
            res = append(res, a)
        }
    }
    
    sort.Ints(res)
    
    return res
}