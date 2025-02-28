type key struct {
    p pos
    min int
}

type pos struct {
    i, j int
}

type KeyHeap []key

func (h KeyHeap) Len() int           { return len(h) }
func (h KeyHeap) Less(i, j int) bool { return h[i].min > h[j].min } // max of minimum
func (h KeyHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *KeyHeap) Push(x interface{}) {
    *h = append(*h, x.(key))
}

func (h *KeyHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maximumMinimumPath(grid [][]int) int {
    h := &KeyHeap{ key{ pos{ 0, 0 }, grid[0][0] } }
    heap.Init(h)
    n := len(grid)
    m := len(grid[0])
    visited := make(map[pos] bool)
    
    helper := func(i, j, min int) {
        p := pos{ i, j }
        if i >= 0 && j >= 0 && i < n && j < m && !visited[p] {
            if grid[i][j] < min {
                min = grid[i][j]
            }
            heap.Push(h, key{ p, min })
            visited[p] = true
        }
    }
    
    for h.Len() != 0 {
        pop := heap.Pop(h).(key)
        
        if pop.p.i == n - 1 && pop.p.j == m - 1 {
            return pop.min
        }
        
        helper(pop.p.i + 1, pop.p.j, pop.min)
        helper(pop.p.i - 1, pop.p.j, pop.min)
        helper(pop.p.i, pop.p.j + 1, pop.min)
        helper(pop.p.i, pop.p.j - 1, pop.min)
    }
    return -1
}