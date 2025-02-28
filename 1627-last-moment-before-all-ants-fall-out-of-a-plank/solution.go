func getLastMoment(n int, left []int, right []int) int {
    max := 0

    for i := 0; i < len(right); i++ {
        t := n - right[i]
        if t > max {
            max = t
        }
    }
    for i := 0; i < len(left); i++ {
        t := left[i]
        if t > max {
            max = t
        }
    }

    return max
}