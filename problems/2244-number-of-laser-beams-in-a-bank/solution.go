func numberOfBeams(bank []string) int {
    prevNumDevices := 0
    res := 0
    
    for i := 0; i < len(bank); i++ {
        counter := strings.Count(bank[i], "1")
        
        res += prevNumDevices * counter
        if counter != 0 {
            prevNumDevices = counter
        }
    }
    return res
}