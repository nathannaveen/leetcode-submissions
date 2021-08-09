type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
    x := old[ len(old) - 1 ]
    *h = old[0 : len(old) - 1 ]
	return x
}

type SeatManager struct {
    h *IntHeap
}


func Constructor(n int) SeatManager {
    h := &IntHeap{}
    for i := 1; i <= n; i++ {
        heap.Push(h, i)
    }
    return SeatManager{ h }
}


func (this *SeatManager) Reserve() int {
    return heap.Pop(this.h).(int)
}


func (this *SeatManager) Unreserve(seatNumber int)  {
    heap.Push(this.h, seatNumber)
}


/**
 * Your SeatManager object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Reserve();
 * obj.Unreserve(seatNumber);
 */