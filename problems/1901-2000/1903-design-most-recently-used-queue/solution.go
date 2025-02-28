type MRUQueue struct {
    queue []int
}


func Constructor(n int) MRUQueue {
    var queue []int

    for i := 1; i <= n; i++ {
        queue = append(queue, i)
    }
    return MRUQueue{queue}
}


func (this *MRUQueue) Fetch(k int) int {
    x := this.queue[k - 1]
    this.queue = append(this.queue[:k - 1], append(this.queue[k:], this.queue[k - 1])...)
    return x
}


/**
 * Your MRUQueue object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Fetch(k);
 */
