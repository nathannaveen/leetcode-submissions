func findLongestChain(pairs [][]int) int {
    sort.Slice(pairs, func(i, j int) bool {
        return pairs[i][1] < pairs[j][1]
    })
    
    counter := 1
    previous := pairs[0][1]
    
    for i := 1; i < len(pairs); i++ {
        if previous < pairs[i][0] {
            previous = pairs[i][1]
            counter++
        }
    }

    return counter
}