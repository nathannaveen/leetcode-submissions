type HitCounter struct {
    arr []int
}


func Constructor() HitCounter {
    return HitCounter{ []int{} }
}


func (this *HitCounter) Hit(timestamp int)  {
    this.helper(timestamp)
    this.arr = append(this.arr, timestamp)
}


func (this *HitCounter) GetHits(timestamp int) int {
    this.helper(timestamp)
    return len(this.arr)
}

func (this *HitCounter) helper(timestamp int) {
    for len(this.arr) > 0 && this.arr[0] <= timestamp - 300 {
        this.arr = this.arr[1:]
    }
}

/**
 * Your HitCounter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Hit(timestamp);
 * param_2 := obj.GetHits(timestamp);
 */