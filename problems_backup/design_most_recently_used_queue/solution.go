type MRUQueue struct {
    arr []int
}

func Constructor(n int) MRUQueue {
    arr := make([]int, n)
    for i := 0; i < n; i++ {
        arr[i] = i + 1
    }
    return MRUQueue{ arr }
}


func (this *MRUQueue) Fetch(k int) int {
    this.arr = append(this.arr, this.arr[k - 1])
    this.arr = append(this.arr[:k - 1], this.arr[k:]...)
    return this.arr[len(this.arr) - 1]
}


/**
 * Your MRUQueue object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Fetch(k);
 */