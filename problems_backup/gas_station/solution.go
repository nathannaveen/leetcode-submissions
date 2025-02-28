func canCompleteCircuit(gas []int, cost []int) int {
    start := 0
    sum := 0
    totalSum := 0
    
    for i := 0; i < len(gas); i++ {
        sum += gas[i] - cost[i]
        
        if sum < 0 {
            totalSum += sum
            sum = 0
            start = i + 1
        }
    }
    totalSum += sum
    
    if totalSum < 0 {
        return -1
    }
    return start
}