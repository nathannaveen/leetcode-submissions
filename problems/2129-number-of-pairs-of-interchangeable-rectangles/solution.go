func interchangeableRectangles(rectangles [][]int) int64 {
    m := make(map[float64] int) // val, counter
    
    res := int64(0)
    
    for _, rectangle := range rectangles {
        temp := float64(rectangle[0]) / float64(rectangle[1])
        res += int64(m[temp])
        m[temp]++
    }
    
    return res
}