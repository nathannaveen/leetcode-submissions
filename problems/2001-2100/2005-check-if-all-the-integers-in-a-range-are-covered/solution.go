func isCovered(ranges [][]int, left int, right int) bool {
    sort.Slice(ranges, func(i, j int) bool {
        return ranges[i][0] < ranges[j][0]
    })
    
    for _, r := range ranges {
        if r[0] > left {
            return false
        } else if r[1] + 1 > left {
            if r[1] < right {
                left = r[1] + 1
            } else {
                return true
            }
        }
        
    }
    
    return false
}