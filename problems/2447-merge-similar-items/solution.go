func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
    items1 = append(items1, items2...)
    
    sort.Slice(items1, func(i, j int) bool {
        return items1[i][0] < items1[j][0] 
    })
    
    i := 0
    
    res := [][]int{}
    
    for i < len(items1) {
        if i < len(items1) {
            if len(res) != 0 && items1[i][0] == res[len(res) - 1][0] {
                res[len(res) - 1][1] += items1[i][1]
            } else {
                res = append(res, items1[i])
            }
            i++
        }
    }
    
    return res
}