type KeyHeap []key

func (h KeyHeap) Len() int           { return len(h) }
func (h KeyHeap) Less(i, j int) bool { return h[i].dist < h[j].dist }
func (h KeyHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *KeyHeap) Push(x interface{}) { *h = append(*h, x.(key)) }

func (h *KeyHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type pos struct {
    i, j int
}

type key struct {
    p pos // position of current move
    dist int // distance from enterence
}



func nearestExit(maze [][]byte, entrance []int) int {
    visited := make(map[pos] bool)
    visited[pos{ entrance[0], entrance[1] }] = true
    h := &KeyHeap{ key{ pos{ entrance[0], entrance[1] }, 0 } }
    n, m := len(maze), len(maze[0])
    
    addToHeap := func(i, j, dist int) {
        if i >= 0 && j >= 0 && i < n && j < m && !visited[pos{ i, j }] && maze[i][j] != '+' {
            heap.Push(h, key{ pos{ i, j }, dist })
            visited[pos{ i, j }] = true
        }
    }
    
    for h.Len() > 0 {
        pop := heap.Pop(h).(key)
        i, j, dist := pop.p.i, pop.p.j, pop.dist
        
        if (i == 0 || j == 0 || i == n - 1 || j == m - 1) && dist != 0 {
            // if position is on the edge, and position is not where we started.
            return dist
        }
        
        addToHeap(i + 1, j, dist + 1)
        addToHeap(i - 1, j, dist + 1)
        addToHeap(i, j + 1, dist + 1)
        addToHeap(i, j - 1, dist + 1)
    }
    
    return -1
}

