type CustomStack struct {
    stack []int
    max int
}


func Constructor(maxSize int) CustomStack {
    return CustomStack{[]int{}, maxSize}
}


func (this *CustomStack) Push(x int)  {
    if this.max > len(this.stack) {
        this.stack = append(this.stack, x)
    }
}


func (this *CustomStack) Pop() int {
    if len(this.stack) == 0 { return -1 }
    temp := this.stack[len(this.stack) - 1]
    this.stack = this.stack[:len(this.stack) - 1]
    return temp
}


func (this *CustomStack) Increment(k int, val int)  {
    for i := 0; i < int(math.Min(float64(len(this.stack)), float64(k))); i++ {
        this.stack[i] += val
    }
}


/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */