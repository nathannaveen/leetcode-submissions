type point struct {
    i, j int
    maxDiff int
}

type PointHeap []point

func (h PointHeap) Len() int           { return len(h) }
func (h PointHeap) Less(i, j int) bool { return h[i].maxDiff < h[j].maxDiff }
func (h PointHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *PointHeap) Push(x interface{}) {
	*h = append(*h, x.(point))
}

func (h *PointHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minimumEffortPath(heights [][]int) int {
    h := &PointHeap{ point{ 0, 0, 0 } }
    heap.Init(h)
    
    visited := map[point] bool {}
    visited[point{ 0, 0, 0 }] = true
    
    helper := func(i, j, prevHeight, maxDiff int) {
        if i < 0 || j < 0 || i >= len(heights) || j >= len(heights[0]) || visited[point{ i, j, 0 }] {
            return
        }
        
        if int(math.Abs(float64(heights[i][j] - prevHeight))) > maxDiff {
            maxDiff = int(math.Abs(float64(heights[i][j] - prevHeight)))
        }
        heap.Push(h, point{ i, j, maxDiff })
    }
    
    for h.Len() != 0 {
        pop := heap.Pop(h).(point)
        visited[point{ pop.i, pop.j, 0 }] = true
        
        if pop.i == len(heights) - 1 && pop.j == len(heights[0]) - 1 {
            return pop.maxDiff
        }
        
        helper(pop.i + 1, pop.j, heights[pop.i][pop.j], pop.maxDiff)
        helper(pop.i - 1, pop.j, heights[pop.i][pop.j], pop.maxDiff)
        helper(pop.i, pop.j + 1, heights[pop.i][pop.j], pop.maxDiff)
        helper(pop.i, pop.j - 1, heights[pop.i][pop.j], pop.maxDiff)
    }
    
    return -1
}



