func minimumLevels(possible []int) int {
    sum := 0

    for _, p := range possible {
        if p == 1 {
            sum++
        } else {
            sum--
        }
    }

    sum2 := 0

    for i := 0; i < len(possible) - 1; i++ {
        if possible[i] == 1 {
            sum2++
        } else {
            sum2--
        }

        if sum2 > sum - sum2 {
            return i + 1
        }
    }

    return -1
}
