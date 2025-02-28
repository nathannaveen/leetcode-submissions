func subsets(nums []int) [][]int {
    return helper([]int{}, nums)
}

func helper(arr, nums []int) [][]int {
    res := [][]int{ arr }
    for i := 0; i < len(nums); i++ {
        newArr := make([]int, len(arr))
        copy(newArr, arr)
        newArr = append(newArr, nums[i])
        res = append(res, helper(newArr, nums[i + 1:])...)
    }
    return res
}