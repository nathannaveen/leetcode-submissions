func garbageCollection(garbage []string, travel []int) int {
    last := make([]int, 3) 
    
    index := map[rune] int { 'M' : 0, 'P' : 1, 'G' : 2 }
    
    prefix := []int{0}
    
    res := 0
    
    for _, t := range travel {
        prefix = append(prefix, prefix[len(prefix) - 1] + t)
    }
    
    for i, g := range garbage {
        for _, l := range g {
            z := index[l]
            
            res += prefix[i] - prefix[last[z]] + 1
            last[z] = i
        }
    }
    
    return res
}