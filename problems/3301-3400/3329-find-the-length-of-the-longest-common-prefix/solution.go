func longestCommonPrefix(arr1 []int, arr2 []int) int {
    last := 0
    m := map[int] bool {}
    
    for _, val := range arr1 {
        n := val
        for n > 0 {
            m[n] = true
            n /= 10
        }
    }
    
    for _, val := range arr2 {
        n := val
        for n > 0 {
            if m[n] && n > last {
                last = n
            }
            n /= 10
        }
    }
    
    res := 0
    
    for last > 0 {
        last /= 10
        res++
    }
    
    return res
}