type RecentCounter struct {
    k int
    arr []int
}


func Constructor() RecentCounter {
    return RecentCounter{0, []int{}}
}


func (this *RecentCounter) Ping(t int) int {
    this.arr = append(this.arr, t)
    
    for this.arr[this.k] < t - 3000 {
        this.k++
    }
    return len(this.arr) - this.k
}


/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */