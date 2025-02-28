func matchPlayersAndTrainers(players []int, trainers []int) int {
    sort.Ints(players)
    sort.Ints(trainers)
    
    j := 0
    
    res := 0
    
    for i := 0; i < len(players); i++ {
        for j < len(trainers) - 1 && players[i] > trainers[j] {
            j++
        }
        
        if players[i] <= trainers[j] {
            res++
            j++
        }
        
        if j == len(trainers) {
            break
        }
    }
    
    return res
}