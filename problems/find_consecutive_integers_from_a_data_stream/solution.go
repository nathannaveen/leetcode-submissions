type DataStream struct {
    m map[int] int // map[index] val
    i int // index for m
    n int // number of last k values not equal to value
    k int
    value int
}


func Constructor(value int, k int) DataStream {
    return DataStream{map[int] int {}, 0, 0, k, value}
}


func (this *DataStream) Consec(num int) bool {
    if num != this.value {
        this.n++
    }
    if this.i >= this.k && this.m[this.i - this.k] != this.value {
        this.n--
    }
    this.m[this.i] = num
    this.i++

    if this.i < this.k {
        return false
    }

    return this.n == 0
}


/**
 * Your DataStream object will be instantiated and called as such:
 * obj := Constructor(value, k);
 * param_1 := obj.Consec(num);
 */