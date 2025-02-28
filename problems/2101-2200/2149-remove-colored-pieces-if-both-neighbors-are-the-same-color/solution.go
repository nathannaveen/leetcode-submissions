func winnerOfGame(colors string) bool {
    aCounter, bCounter := 0, 0
    prev := 'a'
    counter := 0
    
    for _, color := range colors {
        if color == prev {
            counter++
        } else {
            counter = 0
        }
        if counter >= 2 && color == prev {
            if color == 'A' {
                aCounter++
                bCounter--
            }
            bCounter++
        }
        prev = color
    }
    
    return aCounter > bCounter
}