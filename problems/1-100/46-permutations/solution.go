func permute(nums []int) [][]int {
    return helper(nums, []int{})
}

func helper(nums, arr []int) [][]int {
    if len(nums) == 0 {
        return [][]int{ arr }
    }
    
    res := [][]int{}
    
    for i, n := range nums {
        newNums := make([]int, len(nums))
        copy(newNums, nums)
        newArr := make([]int, len(arr))
        copy(newArr, arr)
        
        res = append(res, helper(append(newNums[:i], newNums[i + 1:]...), append(newArr, n))...)
    }
    
    return res
}