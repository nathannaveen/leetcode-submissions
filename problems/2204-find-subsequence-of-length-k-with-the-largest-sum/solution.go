type key struct {
    i int // index
    val int
}

func maxSubsequence(nums []int, k int) []int {
    arr := []key{}
    res := []int{}
    
    for i := 0; i < len(nums); i++ {
        arr = append(arr, key{ i, nums[i] })
    }
    
    sort.Slice(arr, func(i, j int) bool {
        return arr[i].val > arr[j].val
    })
    
    arr = arr[: k]
    
    sort.Slice(arr, func(i, j int) bool {
        return arr[i].i < arr[j].i
    })
    
    for _, p := range arr {
        res = append(res, p.val)
    }
    
    return res
}