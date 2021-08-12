func findLonelyPixel(picture [][]byte) int {
    cols := make([]int, len(picture[0]))
    rows := make([]int, len(picture))
    res := 0
    
    for i := range picture {
        for j := range picture[0] {
            if picture[i][j] == 'B' {
                rows[i]++
                cols[j]++
            }
        }
    }
    
    for i := range picture {
        for j := range picture[0] {
            if picture[i][j] == 'B' && rows[i] == 1 && cols[j] == 1 {
                res++
            }
        }
    }
    
    return res
}