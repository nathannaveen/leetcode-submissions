type SnakeGame struct {
    counter int
    height int
    width int
    food [][]int
    foodCounter int
    m map[point] bool
    previous []point
}

type point struct {
    i, j int
}


func Constructor(width int, height int, food [][]int) SnakeGame {
    m := make(map[point] bool)
    m[point{ 0, 0 }] = true
    return SnakeGame{ 0, height, width, food, 0, m, []point{ point{0, 0} } }
}


func (this *SnakeGame) Move(direction string) int {
    newPoint := this.previous[len(this.previous) - 1]
    
    switch direction {
        case "R":
            newPoint.j++
        case "L":
            newPoint.j--
        case "U":
            newPoint.i--
        case "D":
            newPoint.i++
    }
    
    if newPoint.i < 0 || newPoint.j < 0 || newPoint.i >= this.height || newPoint.j >= this.width { 
        // If we run into walls
        return -1 
    }
    
    if this.foodCounter < len(this.food) && this.food[this.foodCounter][0] == newPoint.i && this.food[this.foodCounter][1] == newPoint.j {
        // If on a piece of food
        this.foodCounter++
        this.counter++
        
    } else {
        // If not on food we can remove the end of the tail
        pop := this.previous[0]
        this.previous = this.previous[1:]
        delete(this.m, pop)
        
    }
    
    if this.m[newPoint] { 
        // Running into itself
        return -1 
    }
    
    this.previous = append(this.previous, newPoint)
    
    this.m[newPoint] = true
    
    return this.counter
}


/**
 * Your SnakeGame object will be instantiated and called as such:
 * obj := Constructor(width, height, food);
 * param_1 := obj.Move(direction);
 */