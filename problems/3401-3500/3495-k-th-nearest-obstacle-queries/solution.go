import "container/heap"

func resultsArray(queries [][]int, k int) []int {
    h := &IntHeap{}
    heap.Init(h)
    res := make([]int, len(queries))

    for i, q := range queries {
        val := abs(q[0]) + abs(q[1])
        if h.Len() >= k {
            if (*h)[0] > val {
                heap.Pop(h)
                heap.Push(h, val)
            }
        } else {
            heap.Push(h, val)
        }
        
        if h.Len() < k {
            res[i] = -1
        } else {
            res[i] = (*h)[0]
        }
    }

    return res
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}
