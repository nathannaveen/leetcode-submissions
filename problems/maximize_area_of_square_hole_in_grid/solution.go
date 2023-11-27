func maximizeSquareHoleArea(n int, m int, hBars []int, vBars []int) int {
    sort.Ints(vBars)
    sort.Ints(hBars)

    a := findContinuous(hBars)
    b := findContinuous(vBars)
    return min(a * a, b * b)
}

func findContinuous(bars []int) int {
    continuous := 0
    maxContinuous := 0
    prevBar := 0

    for _, bar := range bars {
        if bar != prevBar + 1 {
            if continuous > maxContinuous {
                maxContinuous = continuous
            }
            continuous = 1
        } else {
            continuous++
        }
        prevBar = bar
    }
    if continuous > maxContinuous {
        maxContinuous = continuous
    }
    
    return maxContinuous + 1
}

func min(a, b int) int {
    if a > b {
        return b
    }

    return a
}