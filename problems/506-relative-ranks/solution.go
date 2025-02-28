type key struct {
    val int
    index int
}

func findRelativeRanks(score []int) []string {
    arr := []key{}
    res := make([]string, len(score))
    
    for i, s := range score {
        arr = append(arr, key{ s, i })
    }
    
    sort.Slice(arr, func(i, j int) bool {
        return arr[i].val > arr[j].val
    })
    
    for i := 0; i < len(arr); i++ {
        if i == 0 {
            res[arr[i].index] = "Gold Medal"
        } else if i == 1 {
            res[arr[i].index] = "Silver Medal"
        } else if i == 2 {
            res[arr[i].index] = "Bronze Medal"
        } else {
            res[arr[i].index] = strconv.Itoa(i + 1)
        }
    }
    
    return res
}