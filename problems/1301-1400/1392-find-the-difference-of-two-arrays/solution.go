func findDifference(nums1 []int, nums2 []int) [][]int {
    res := [][]int{ {}, {}}
    
    one := map[int] bool {}
    oneFound := map[int] bool {}
    
    two := map[int] bool {}
    twoFound := map[int] bool {}
    
    for _, n := range nums1 {
        one[n] = true
    }
    
    for _, n := range nums2 {
        if !one[n] && !twoFound[n] {
            res[1] = append(res[1], n)
        }
        two[n] = true
        twoFound[n] = true
    }
    
    for _, n := range nums1 {
        if !two[n] && !oneFound[n] {
            res[0] = append(res[0], n)
        }
        oneFound[n] = true
    }
    
    return res
}