func areaOfMaxDiagonal(dimensions [][]int) int {
    x := 0
    for i, d := range dimensions {
        z := math.Sqrt(float64(d[0] * d[0] + d[1] * d[1]))
        z2 := math.Sqrt(float64(dimensions[x][0] * dimensions[x][0] + dimensions[x][1] * dimensions[x][1]))
        
        k := d[0] * d[1]
        k2 := dimensions[x][0] * dimensions[x][1]
        
        if z > z2 || (z == z2 && k > k2) {
            x = i
        }
    }
    
    return dimensions[x][0] * dimensions[x][1]
}