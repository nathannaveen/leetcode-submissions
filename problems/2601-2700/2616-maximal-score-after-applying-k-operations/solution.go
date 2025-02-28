func maxKelements(nums []int, k int) int64 {
    h := &IntHeap{}
    heap.Init(h)
    
    res := int64(0)

    for _, num := range nums {
        heap.Push(h, num)
    }

    for i := 0; i < k; i++ {
        pop := heap.Pop(h).(int)

        res += int64(pop)
        heap.Push(h, int(math.Ceil(float64(pop) / float64(3))))
    }

    return res
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}