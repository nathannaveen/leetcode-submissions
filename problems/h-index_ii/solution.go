func hIndex(citations []int) int {
    for i := 0; i < len(citations); i++ {
        if citations[i] >= len(citations) - i {
            return len(citations) - i
        }
    }
    
    return 0
}