func maximumUnits(boxTypes [][]int, truckSize int) int {
    sort.Slice(boxTypes, func(i, j int) bool {
        return boxTypes[i][1] > boxTypes[j][1]
    })
    
    res := 0
    
    for i := 0; i < len(boxTypes); i++ {
        b := boxTypes[i]
        
        lastBoxes := float64(truckSize * b[1])
        allBoxesOfI := float64(b[0] * b[1])
        
        res += int(math.Max(math.Min(lastBoxes, allBoxesOfI), float64(0)))
        truckSize -= b[0]
    }
    
    return res
}