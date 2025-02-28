func maxUncrossedLines(nums1 []int, nums2 []int) int {
    n := len(nums1)
    m := len(nums2)
    
    arr := make([][]int, n + 1)
    
    for i := 0; i < n + 1; i++ {
        arr[i] = make([]int, m + 1)
        
        for j := 0; j < m + 1; j++ {
            if i != 0 && j != 0 {
                if nums1[i - 1] == nums2[j - 1] {
                    arr[i][j] = 1 + arr[i - 1][j - 1]
                } else {
                    arr[i][j] = int(math.Max(float64(arr[i - 1][j]), float64(arr[i][j - 1])))
                }
            }
        }
    }
    
    return arr[n][m]
}