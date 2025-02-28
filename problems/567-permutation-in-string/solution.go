func checkInclusion(s1 string, s2 string) bool {
    arr := make([]int, 26)
    arr2 := make([]int, 26)
    
    if len(s1) > len(s2) {
        return false
    }
    
    for i := 0; i < len(s1); i++ {
        arr[int(s1[i] - 'a')]++
        arr2[int(s2[i] - 'a')]++
    }
    
    if reflect.DeepEqual(arr, arr2) {
        return true
    }
    
    for i := len(s1); i < len(s2); i++ {
        arr2[int(s2[i] - 'a')]++
        arr2[int(s2[i - len(s1)] - 'a')]--
        
        if reflect.DeepEqual(arr, arr2) {
            return true
        }
    }
    
    return false
}