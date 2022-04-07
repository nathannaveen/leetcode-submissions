func lastStoneWeight(stones []int) int {
    h := &IntHeap{}
	heap.Init(h)
    
    for _, s := range stones {
        heap.Push(h, s)
    }
    
    for h.Len() > 1 {
        heap.Push(h, heap.Pop(h).(int) - heap.Pop(h).(int))
    }
    
    if h.Len() == 0 {
        return 0
    }
    
    return heap.Pop(h).(int)
}

// Heaps structure below:

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
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