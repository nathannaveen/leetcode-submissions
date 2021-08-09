type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	x := old[len(old) - 1]
	*h = old[0 : len(old) - 1]
	return x
}

func minStoneSum(piles []int, k int) int {
    h := &IntHeap{}
    heap.Init(h)
    sum := 0
    
    for _, pile := range piles {
        heap.Push(h, pile)
        sum += pile
    }
    
    for i := 0; i < k; i++ {
        temp := heap.Pop(h).(int)
        sum -= (temp / 2)
        heap.Push(h, temp - (temp / 2))
    }
    
    return sum
}