func minimumBoxes(n int) int {
    base := 0
    total := 0
    add := 0

    for total < n {
        add++
        base += add
        total += base
    }

    newBase := base - add
    newTotal := total - base

    for newTotal < n {
        newBase++
        newTotal += newBase - base + add
    }

    if total == n {
        return base
    }

    return newBase
}