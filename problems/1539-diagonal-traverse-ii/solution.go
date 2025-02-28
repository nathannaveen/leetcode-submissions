func findDiagonalOrder(nums [][]int) []int {
    arr := [][]int{}
    res := []int{}
    
    for i := 0; i < len(nums); i++ {
        for j := 0; j < len(nums[i]); j++ {
            if len(arr) <= i + j {
                arr = append(arr, []int{})
            }
            arr[i + j] = append(arr[i + j][:0], append([]int{nums[i][j]}, arr[i + j][0:]...)...)
        }
    }
    
    for i := 0; i < len(arr); i++ {
        res = append(res, arr[i]...)
    }
    
    return res
}