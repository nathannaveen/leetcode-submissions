func losingPlayer(x int, y int) string {
    p := true
    
    for x > 0 || y > 0 {
        if x - 1 < 0 || y - 4 < 0 {
            if !p {
                return "Alice"
            } else {
                return "Bob"
            }
        }
        x -= 1
        y -= 4
        p = !p
    }
    
    if !p {
        return "Alice"
    } else {
        return "Bob"
    }
}
