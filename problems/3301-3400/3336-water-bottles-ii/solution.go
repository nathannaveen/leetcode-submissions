func maxBottlesDrunk(numBottles int, numExchange int) int {
    res := 0
    done := false
    empty := 0

    for !done {
        if numBottles > 0 {
            empty += numBottles
            res += numBottles
            numBottles = 0
            continue
        }
        if empty >= numExchange {
            empty -= numExchange
            numBottles++
            numExchange++
            continue
        }

        done = true
    }

    return res
}
