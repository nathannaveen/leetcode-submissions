func minimumTime(jobs []int, workers []int) int {
    sort.Ints(jobs)
    sort.Ints(workers)
    
    res := 0
    
    for i := 0; i < len(jobs); i++ {
        n := int(math.Ceil(float64(jobs[i]) / float64(workers[i])))
        
        if n > res {
            res = n
        }
    }
    
    return res
}
