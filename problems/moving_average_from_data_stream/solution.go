type MovingAverage struct {
    arr []int
    sum int
    size int
}


/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
    return MovingAverage{[]int{}, 0, size}
}


func (this *MovingAverage) Next(val int) float64 {
    this.arr = append(this.arr, val)
    this.sum += val
    
    if this.size < len(this.arr) {
        this.sum -= this.arr[0]
        this.arr = this.arr[1:]
    }
    
    return float64(this.sum) / float64(len(this.arr))
}


/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */