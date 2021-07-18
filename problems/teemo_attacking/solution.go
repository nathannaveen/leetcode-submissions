func findPoisonedDuration(timeSeries []int, duration int) int {
    if len(timeSeries) == 1 {
        return duration
    }
    
    timeSeries[0] = int(math.Min(float64(duration), float64(timeSeries[1] - timeSeries[0])))
    
    for i := 1; i < len(timeSeries) - 1; i++ {
        timeSeries[i] = int(math.Min(float64(timeSeries[i - 1] + duration), float64(timeSeries[i - 1] + timeSeries[i + 1] - timeSeries[i])))
    }
    
    return timeSeries[len(timeSeries) - 2] + duration
}