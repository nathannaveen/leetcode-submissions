func divideArray(nums []int, k int) [][]int {
    sort.Ints(nums)
    res := [][]int{}
    
    for i := 0; i < len(nums); i += 3 {
        if nums[i + 2] - nums[i] <= k {
            res = append(res, nums[i : i + 3])
        } else {
            return [][]int{}
        }
    }
    
    return res
}
