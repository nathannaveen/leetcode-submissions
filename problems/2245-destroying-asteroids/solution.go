func asteroidsDestroyed(mass int, asteroids []int) bool {
    sort.Ints(asteroids)
    
    for _, asteroid :=  range asteroids {
        if asteroid > mass {
            return false
        }
        mass += asteroid
    }
    
    return true
}