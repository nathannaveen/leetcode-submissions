func minMeetingRooms(intervals [][]int) int {
    res := 0
    
    s, e := make([]int, len(intervals)), make([]int, len(intervals))
    
    for i, interval := range intervals {
        s[i] = interval[0]
        e[i] = interval[1]
    }
    
    sort.Ints(s)
    sort.Ints(e)
    
    j := 0
    
    for _, startVal := range s {
        if startVal < e[j] {
            res++
            j--
        }
        j++
    }
    
    return res
}