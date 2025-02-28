func minimumPerimeter(neededApples int64) int64 {
    sideLength := 0
    
    for neededApples > int64(0) {
        sideLength++
        neededApples -= int64(sideLength * sideLength * 12)
    }
    
    return int64(sideLength * 8)
}