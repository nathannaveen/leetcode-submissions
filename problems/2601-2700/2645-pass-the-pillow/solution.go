func passThePillow(n int, time int) int {
    time = time % (2 * (n - 1)) + 1
    return int(math.Min(float64(time), float64(n * 2 - time)))
}