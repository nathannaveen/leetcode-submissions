func maxSatisfaction(satisfaction []int) int {
    sort.Ints(satisfaction)
    max := float64(0)
    
    for i := 0; i < len(satisfaction); i++ {
        sum := 0
        for j := i; j < len(satisfaction); j++ {
            sum += (j - i + 1) * satisfaction[j]
        }
        max = math.Max(float64(sum), max)
    }
    return int(max)
}