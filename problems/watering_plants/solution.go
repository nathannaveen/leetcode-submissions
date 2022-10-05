func wateringPlants(plants []int, capacity int) int {
    can := capacity
    res := 0
    cur := -1
    
    for i := 0; i < len(plants) - 1; i++ {
        res += i - cur
        can -= plants[i]
        
        if plants[i + 1] > can {
            can = capacity
            res += i + 1
            cur = -1
        } else {
            cur = i
        }
    }
    
    res += len(plants) - 1 - cur
    
    return res
}