func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
    res := []int{}
    arr := [][]int{ make([]int, 101), make([]int, 101), make([]int, 101) }
    
    for _, n := range nums1 { arr[0][n] = 1 }
    for _, n := range nums2 { arr[1][n] = 1 }
    for _, n := range nums3 { arr[2][n] = 1 }
    
    for i := 0; i < 100; i++ {
        if arr[0][i + 1] + arr[1][i + 1] + arr[2][i + 1] >= 2 {
            res = append(res, i + 1)
        }
    }
    
    return res
}