func countTestedDevices(batteryPercentages []int) int {
    x := 0
    res := 0

    for _, b := range batteryPercentages {
        if b - x > 0 {
            res++
            x++
        }
    }

    return res
}