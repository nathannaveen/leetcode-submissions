type key struct {
    PT int // plant time
    GT int // grow time
}

func earliestFullBloom(plantTime []int, growTime []int) int {
    time := 0
    flowers := []key{}
    max := 0
    
    for i := 0; i < len(plantTime); i++ {
        flowers = append(flowers, key{ plantTime[i], growTime[i] })
    }
    
    sort.Slice(flowers, func(i, j int) bool { return flowers[i].GT > flowers[j].GT })
    
    for i := 0; i < len(flowers); i++ {
        time += flowers[i].PT
        
        if time + flowers[i].GT > max {
            max = time + flowers[i].GT
        }
    }
    
    return max
}