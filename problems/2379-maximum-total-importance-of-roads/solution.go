func maximumImportance(n int, roads [][]int) int64 {
    cities := make([]int, n)
    
    for _, road := range roads {
        cities[road[0]]++
        cities[road[1]]++
    }
    
    sort.Slice(cities, func(i, j int) bool {
        return cities[i] > cities[j]
    })
    
    res := int64(0)
    
    for _, city := range cities {
        if city == 0 {
            break
        }
        res += int64(city) * int64(n)
        n--
    }
    
    return res
}