func pivotArray(nums []int, pivot int) []int {
    arr := []int{}
    
    for _, n := range nums {
        if n < pivot { arr = append(arr, n) }
    }
    
    for _, n := range nums {
        if n == pivot { arr = append(arr, n) }
    }
    
    for _, n := range nums {
        if n > pivot { arr = append(arr, n) }
    }
    
    return arr
}