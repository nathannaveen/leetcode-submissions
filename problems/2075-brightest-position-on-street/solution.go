type key struct {
    index int
    val int
}

func brightestPosition(lights [][]int) int {
    m := make(map[int] int) // pos, arr's index
    arr := []key{}
    
    for _, light := range lights {
        m[light[0] - light[1]]++
        m[light[0] + light[1] + 1]--
    }
    
    for a, b := range m {
        arr = append(arr, key{ a, b } )
    }
    
    sort.Slice(arr, func(i, j int) bool {
        return arr[i].index < arr[j].index
    })
    
    sum := 0
    max := 0
    index := 0
    
    for _, a := range arr {
        sum += a.val
        
        if sum > max {
            max = sum
            index = a.index
        }
    }
    
    return index
}
