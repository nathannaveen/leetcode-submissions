func maxPoints(points [][]int) int {
    max := 0
    
    for i := 0; i < len(points); i++ {
        m := make(map[float64] int) // slope : number of points on the slope
        for j := 0; j < len(points); j++ {
            x1, x2, y1, y2 := points[i][0], points[j][0], points[i][1], points[j][1]
            
            slope := float64(y2 - y1) / float64(x2 - x1)
            m[slope]++
        }
        
        for _, b := range m {
            if b > max { max = b }
        }
    }
    
    if len(points) == 1 { return 1 }
    
    return max + 1
}