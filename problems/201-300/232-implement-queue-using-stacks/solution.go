type MyQueue struct {
    arr []int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
    return MyQueue{ []int{} }
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
    this.arr = append(this.arr[:0], append([]int{x}, this.arr[0:]...)...)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
    temp := this.arr[len(this.arr) - 1]
    this.arr = this.arr[:len(this.arr) - 1]
    return temp
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
    return this.arr[len(this.arr) - 1]
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
    return len(this.arr) == 0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */