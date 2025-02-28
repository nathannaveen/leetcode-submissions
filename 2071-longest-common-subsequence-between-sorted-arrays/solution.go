func longestCommomSubsequence(arrays [][]int) []int {
    arr := make([]int, len(arrays)) // the counters for each array
    keepGoing := true
    res := []int{}
    
    for keepGoing {
        keepGoing = false
        temp := 0
        max := arrays[0][arr[0]]
        
        for i := 1; i < len(arrays); i++ {
            if arrays[i][arr[i]] > max {
                max = arrays[i][arr[i]]
            }
        }
        for i := 0; i < len(arrays); i++ {
            if arrays[i][arr[i]] == max {
                temp++
            } else if arrays[i][arr[i]] < max && arr[i] != len(arrays[i]) - 1 {
                keepGoing = true
                arr[i]++
            }
        }
        if temp == len(arrays) {
            res = append(res, max)
            for i := 0; i < len(arrays); i++ {
                if arr[i] != len(arrays[i]) - 1 {
                    arr[i]++
                    keepGoing = true
                }
            }
        }
    }
    return res
}