type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type KthLargest struct {
    h *IntHeap
    k int
}


func Constructor(k int, nums []int) KthLargest {
    h := &IntHeap{}
    heap.Init(h)
    
    for _, n := range nums {
        heap.Push(h, n)
    }
    
    l := h.Len()
    
    for i := 0; i < l - k; i++ {
        heap.Pop(h)
    }
    return KthLargest { h, k }
}

func (this *KthLargest) Add(val int) int {
    heap.Push(this.h, val)
    
    if this.h.Len() > this.k {
        heap.Pop(this.h)
    }
    res := heap.Pop(this.h).(int)
    heap.Push(this.h, res)
    
    return res
}


/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */