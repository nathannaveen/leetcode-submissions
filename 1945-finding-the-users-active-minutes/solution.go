func findingUsersActiveMinutes(logs [][]int, k int) []int {
    res := make([]int, k)
    arr := []map[int] bool{}
    m := make(map[int] int) // value, index in arr
    
    for _, log := range logs {
        
        if _, ok := m[log[0]]; !ok {
            arr = append(arr, make(map[int] bool))
            m[log[0]] = len(arr)
        } else {
            res[len(arr[m[log[0]] - 1]) - 1]--
        }
        arr[m[log[0]] - 1][log[1]] = true
        res[len(arr[m[log[0]] - 1]) - 1]++
    }
    
    return res
}