type Vector2D struct {
    vec [][]int
    i, j int
}


func Constructor(vec [][]int) Vector2D {
    return Vector2D{vec, 0, -1}
}


func (this *Vector2D) Next() int {
    this.j++
    if len(this.vec[this.i]) == this.j {
        this.i++
        this.j = 0
        for this.i < len(this.vec) {
            if len(this.vec[this.i]) > 0 {
                return this.vec[this.i][this.j]
            }
            this.i++
        }
    }
    
    return this.vec[this.i][this.j]
}


func (this *Vector2D) HasNext() bool {
    if this.i == len(this.vec) {
        return false
    }
    if this.j + 1 >= len(this.vec[this.i]) {
        k := 1
        for this.i + k < len(this.vec) {
            if len(this.vec[this.i + k]) > 0 {
                return true
            }
            k++
        }
        return false
    }
    return true
}


/**
 * Your Vector2D object will be instantiated and called as such:
 * obj := Constructor(vec);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */