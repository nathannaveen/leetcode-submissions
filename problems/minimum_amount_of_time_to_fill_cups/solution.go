func fillCups(amount []int) int {
    h := &IntHeap{}
    heap.Init(h)
    
    for _, a := range amount {
        if a > 0 {
            heap.Push(h, a)
        }
    }
    
    res := 0
    
    for h.Len() > 0 {
        n := h.Len()
        shouldPush := []int{}
        for i := 0; i < int(math.Min(float64(2), float64(n))); i++ {
            p := heap.Pop(h).(int)
            shouldPush = append(shouldPush, p - 1)
        }
        
        for i := 0; i < len(shouldPush); i++ {
            if shouldPush[i] > 0 {
                heap.Push(h, shouldPush[i])
            }
        }
        
        res++
    }
    
    return res
}

type IntHeap []int

func (self IntHeap) Len() int {
	return len(self)
}

func (self IntHeap) Less(a, b int) bool {
	return self[a] > self[b]
}
func (self IntHeap) Swap(a, b int) {
	self[a], self[b] = self[b], self[a]
}
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