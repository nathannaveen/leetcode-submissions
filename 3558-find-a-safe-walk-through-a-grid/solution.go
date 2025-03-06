import "container/heap"

func findSafeWalk(grid [][]int, health int) bool {
    h := &KeyHeap{}
    heap.Init(h)

    heap.Push(h, key{health - grid[0][0], 0, 0})
    visited := map[key] bool { {i : 0, j : 0} : true } // key is just being used for the i and j

    for h.Len() > 0 {
        pop := heap.Pop(h).(key)

        if pop.i == len(grid) - 1 && pop.j == len(grid[0]) - 1 {
            return true
        }

        directions := []key{{i : 0, j : 1}, {i : 1, j : 0}, {i : 0, j : -1}, {i : -1, j : 0}}

        for _, d := range directions {
            i := pop.i + d.i
            j := pop.j + d.j
            if visited[key{i : i, j : j}] || !inBounds(i, j, len(grid), len(grid[0])) {
                continue
            }

            newHealth := pop.health - grid[i][j]

            if newHealth > 0 {
                heap.Push(h, key{newHealth, i, j})
                visited[key{i : i, j : j}] = true
            }
        }
    }

    return false
}

func inBounds(i, j, I, J int) bool {
    return i >= 0 && j >= 0 && i < I && j < J
}

type key struct {
    health int
    i, j int
}

type KeyHeap []key

func (h KeyHeap) Len() int           { return len(h) }
func (h KeyHeap) Less(i, j int) bool { return h[i].health > h[j].health }
func (h KeyHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *KeyHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(key))
}

func (h *KeyHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
