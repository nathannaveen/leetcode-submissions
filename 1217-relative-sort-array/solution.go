func relativeSortArray(arr1 []int, arr2 []int) []int {
    arr := make([]int, len(arr2))
    res := []int{}
    addToEnd := []int{}
    
    for _, i := range arr1 {
        pos := 0
        contains := false
        for j := range arr {
            if arr2[j] == i {
                pos = j
                contains = true
                break
            }
        }
        
        if contains {
            res = append(res[:arr[pos]], append([]int{i}, res[arr[pos]:]...)...)

            for j := pos; j < len(arr2); j++ {
                arr[j]++
            }
        } else {
            addToEnd = append(addToEnd, i)
        }
    }
    
    sort.Ints(addToEnd)
    res = append(res, addToEnd...)
    
    return res
}