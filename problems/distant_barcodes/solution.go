type key struct {
    val int
    n int
}

type KeyHeap []key

func (h KeyHeap) Len() int           { return len(h) }
func (h KeyHeap) Less(i, j int) bool { return h[i].n > h[j].n }
func (h KeyHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *KeyHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(key))
}

func (h *KeyHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func rearrangeBarcodes(barcodes []int) []int {
    h := &KeyHeap{}
    heap.Init(h)
    m := make(map[int] int)
    
    for _, b := range barcodes {
        m[b]++
    }
    
    for a, b := range m {
        heap.Push(h, key{a, b} )
    }
    
    for i := 0; i < len(barcodes) - 1; i += 2 {
        temp1 := heap.Pop(h).(key)
        temp2 := heap.Pop(h).(key)
        
        barcodes[i] = temp1.val
        barcodes[i + 1] = temp2.val
        
        if temp1.n - 1 != 0 {
            heap.Push(h, key{temp1.val, temp1.n - 1} )
        }
        
        if temp2.n - 1 != 0 {
            heap.Push(h, key{temp2.val, temp2.n - 1} )
        }
    }
    
    if h.Len() != 0 {
        barcodes[len(barcodes) - 1] = heap.Pop(h).(key).val
    }

    return barcodes
}




