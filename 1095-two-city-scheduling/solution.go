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

func twoCitySchedCost(costs [][]int) int {
    h := &IntHeap{}
    heap.Init(h)
    res := 0
    
    for i := 0; i < len(costs); i++ {
        temp := costs[i][1] - costs[i][0]
        
        res += costs[i][0]
        res += temp
        
        if h.Len() == len(costs) / 2 && temp < (*h)[0] {
            res -= heap.Pop(h).(int)
            heap.Push(h, temp)
        } else if h.Len() != len(costs) / 2 {
            heap.Push(h, temp)
        } else {
            res -= temp
        }
    }
    
    return res
}