type MyCircularQueue struct {
    k int
    arr []int
}


func Constructor(k int) MyCircularQueue {
    return MyCircularQueue{ k, []int{} }
}


func (this *MyCircularQueue) EnQueue(value int) bool {
    if len(this.arr) == this.k { return false }
    this.arr = append(this.arr, value)
    return true
}


func (this *MyCircularQueue) DeQueue() bool {
    if len(this.arr) == 0 { return false }
    this.arr = this.arr[1:]
    return true
}


func (this *MyCircularQueue) Front() int {
    if len(this.arr) == 0 { return -1 }
    return this.arr[0]
}


func (this *MyCircularQueue) Rear() int {
    if len(this.arr) == 0 { return -1 }
    return this.arr[len(this.arr) - 1]
}


func (this *MyCircularQueue) IsEmpty() bool {
    if len(this.arr) == 0 { return true }
    return false
}


func (this *MyCircularQueue) IsFull() bool {
    if len(this.arr) == this.k { return true }
    return false
}


/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */