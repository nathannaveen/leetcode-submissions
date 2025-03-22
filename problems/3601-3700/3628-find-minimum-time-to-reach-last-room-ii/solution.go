import "container/heap"

type key struct {
    i, j int
    t int
    x int
}

func minTimeToReach(moveTime [][]int) int {
    h := &KeyHeap{}
    heap.Init(h)
    heap.Push(h, key{0, 0, 0, 2}) // room (0, 0), at time 0

    visited := map[key] bool {} // we are using the struct key as the key in the map to store the rooms i and j

    for h.Len() > 0 {
        pop := heap.Pop(h).(key)

        if pop.i == len(moveTime) - 1 && pop.j == len(moveTime[0]) - 1 {
            return pop.t
        }

        if visited[key{i : pop.i, j : pop.j}] {
            continue
        }
        visited[key{i : pop.i, j : pop.j}] = true

        for _, direction := range []key{{0, 1, 0, 0}, {1, 0, 0, 0}, {0, -1, 0, 0}, {-1, 0, 0, 0}} {
            i, j := pop.i + direction.i, pop.j + direction.j

            x := 1
            if pop.x == 1 {
                x = 2
            }

            if inBounds(i, j, len(moveTime), len(moveTime[0])) {
                t := max(moveTime[i][j], pop.t) + x

                heap.Push(h, key{i, j, t, x})
            }
        }
    }

    return -1
}

func inBounds(i, j, I, J int) bool {
    return i >= 0 && j >= 0 && i < I && j < J
}

type KeyHeap []key

func (h KeyHeap) Len() int           { return len(h) }
func (h KeyHeap) Less(i, j int) bool { return h[i].t < h[j].t }
func (h KeyHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *KeyHeap) Push(x any) { *h = append(*h, x.(key)) }
func (h *KeyHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
