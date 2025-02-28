func numberOfLines(widths []int, s string) []int {
    lines := 1
    pixels := 0
    
    for _, l := range s {
        pixels += widths[l - 97]
        
        if pixels > 100 {
            pixels = widths[l - 97]
            lines++
        }
    }
    
    return []int{lines, pixels}
}