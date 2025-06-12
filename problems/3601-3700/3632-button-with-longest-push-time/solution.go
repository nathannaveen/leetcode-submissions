func buttonWithLongestTime(events [][]int) int {
    index := events[0][0]
    time := events[0][1]
    
    for i := 1; i < len(events); i++ {
        d := events[i][1] - events[i - 1][1]
        if d > time || (d == time && events[i][0] < index) {
            index = events[i][0]
            time = d
        }
    }

    return index
}
