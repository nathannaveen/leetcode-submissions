func distanceTraveled(mainTank int, additionalTank int) int {
    res := 0
    for mainTank >= 5 && additionalTank > 0 {
        res += 50
        mainTank -= 4
        additionalTank--
    }
    res += mainTank * 10
    return res
}