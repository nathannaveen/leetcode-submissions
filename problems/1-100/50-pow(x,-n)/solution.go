func myPow(x float64, n int) float64 {
    if n < 0 {
        return 1 / math.Pow(x, float64(-n))
    }
    return math.Pow(x, float64(n))
}