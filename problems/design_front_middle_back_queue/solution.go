type FrontMiddleBackQueue struct {
    slice []int
}


func Constructor() FrontMiddleBackQueue {
    return FrontMiddleBackQueue{ []int{} }
}


func (this *FrontMiddleBackQueue) PushFront(val int)  {
    this.slice = append(this.slice[:0], append([]int{val}, this.slice[0:]...)...)
}


func (this *FrontMiddleBackQueue) PushMiddle(val int)  {
    this.slice = append(this.slice[: len(this.slice) / 2], append([]int{val}, this.slice[len(this.slice) / 2:]...)...)
}


func (this *FrontMiddleBackQueue) PushBack(val int)  {
    this.slice = append(this.slice[:len(this.slice)], append([]int{val}, this.slice[len(this.slice):]...)...)
}


func (this *FrontMiddleBackQueue) PopFront() int {
    if len(this.slice) == 0 { return -1 }
    res := this.slice[0]
    this.slice = this.slice[1:]
    
    return res
}


func (this *FrontMiddleBackQueue) PopMiddle() int {
    if len(this.slice) == 0 { return -1 }
    res := 0
    if len(this.slice) % 2 == 0 {
        res = this.slice[len(this.slice) / 2 - 1]
        this.slice = append(this.slice[:len(this.slice) / 2 - 1], this.slice[len(this.slice) / 2:]...)
    } else {
        res = this.slice[len(this.slice) / 2]
        this.slice = append(this.slice[:len(this.slice) / 2], this.slice[len(this.slice) / 2 + 1:]...)
    }
    return res
}


func (this *FrontMiddleBackQueue) PopBack() int {
    if len(this.slice) == 0 { return -1 }
    res := this.slice[len(this.slice) - 1]
    this.slice = append(this.slice[:len(this.slice) - 1], this.slice[len(this.slice):]...)
    
    return res
}


/**
 * Your FrontMiddleBackQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PushFront(val);
 * obj.PushMiddle(val);
 * obj.PushBack(val);
 * param_4 := obj.PopFront();
 * param_5 := obj.PopMiddle();
 * param_6 := obj.PopBack();
 */