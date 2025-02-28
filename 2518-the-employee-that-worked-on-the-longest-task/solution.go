func hardestWorker(n int, logs [][]int) int {
    id, max := -1, -1
    time := 0
    
    for _, log := range logs {
        t := log[1] - time
        if (t > max) || (t == max && log[0] < id) {
            id = log[0]
            max = t
        }
        time = log[1]
    }
    
    return id
}