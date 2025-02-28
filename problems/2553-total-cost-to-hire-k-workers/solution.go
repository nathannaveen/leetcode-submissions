type key struct {
    cost int
    index int
    fromRight int // 1 moving from the right end, or 2 the left end
}

func totalCost(costs []int, k int, candidates int) int64 {
    h := &KeyHeap{}
    heap.Init(h)
    n := 0
    leftIndex, rightIndex := 0, len(costs) - 1 
    res := int64(0)

    for i := 0; i < candidates; i++ {
        heap.Push(h, key{costs[i], i, 2})
        leftIndex = i
        n++
        if n == len(costs) {
            break
        }
        heap.Push(h, key{costs[len(costs) - 1 - i], len(costs) - 1 - i, 1})
        rightIndex = len(costs) - 1 - i
        n++
        if n == len(costs) {
            break
        }
    }

    for i := 0; i < k; i++ {
        pop := heap.Pop(h).(key)
        res += int64(pop.cost)

        if n != len(costs) {
            if pop.fromRight == 1 {
                rightIndex--
                heap.Push(h, key{costs[rightIndex], rightIndex, 1})
            } else {
                leftIndex++
                heap.Push(h, key{costs[leftIndex], leftIndex, 2})
            }
            n++
        }
    }

    return res
}

type KeyHeap []key

func (h KeyHeap) Len() int           { return len(h) }
func (h KeyHeap) Less(i, j int) bool { 
    if h[i].cost == h[j].cost {
        return h[i].index < h[j].index
    }
    return h[i].cost < h[j].cost 
}
func (h KeyHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *KeyHeap) Push(x interface{}) { *h = append(*h, x.(key)) }

func (h *KeyHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
