func maximumPoints(enemyEnergies []int, currentEnergy int) int64 {
    sort.Ints(enemyEnergies)
    res := int64(0)
    i := len(enemyEnergies) - 1

    if currentEnergy < enemyEnergies[0] {
        return 0
    }

    for i >= 0 {
        if currentEnergy >= enemyEnergies[0] {
            res += int64(currentEnergy / enemyEnergies[0])
            currentEnergy %= enemyEnergies[0]
        } else {
            currentEnergy += enemyEnergies[i]
            i--
        }
    }

    return res
}
