import (
	"container/heap"
	"fmt"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func connectSticks(sticks []int) int {
    h := &IntHeap{}
    heap.Init(h)
    res := 0
    
    for _, stick := range sticks {
        heap.Push(h, stick)
    }
    
    for h.Len() != 1 {
        temp := heap.Pop(h).(int) + heap.Pop(h).(int)
        res += temp
        heap.Push(h, temp)
    }
    
    return res
}