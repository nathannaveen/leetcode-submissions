func maxDistToClosest(seats []int) int {
    first := -1
    res := 0
    
    for i := range seats {
        if seats[i] == 1 {
            if first == -1 {
                res = i
                first = i
            } else {
                res = int(math.Max(float64(res), float64(i - first) / 2))
                first = i
            }
        }
    }
    
    res = int(math.Max(float64(res), float64(len(seats) - 1 - first)))
    
    return res
}