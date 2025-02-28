type OrderedStream struct {
    i int
    m map[int] string
    n int
}


func Constructor(n int) OrderedStream {
    return OrderedStream{ 1, map[int] string {}, n }
}


func (this *OrderedStream) Insert(idKey int, value string) []string {
    this.m[idKey] = value
    res := []string{}

    for this.i <= this.n {
        if val, ok := this.m[this.i]; ok {
            res = append(res, val)
        } else {
            break
        }
        this.i++
    }

    return res
}


/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */