func trimMean(arr []int) float64 {
    sort.Ints(arr)
    sum := 0
    
    for i := len(arr) / 20; i < len(arr) - (len(arr) / 20); i++ {
        sum += arr[i]
    }
    return float64(sum) / float64(len(arr) - (len(arr) / 20) * 2)
}