type key struct {
    i, j int
    height int
}

type KeyHeap []key

func (h KeyHeap) Len() int           { return len(h) }
func (h KeyHeap) Less(i, j int) bool { return h[i].height < h[j].height }
func (h KeyHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *KeyHeap) Push(x interface{}) { *h = append(*h, x.(key)) }

func (h *KeyHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func trapRainWater(heightMap [][]int) int {
    h := &KeyHeap{}
    heap.Init(h)
    n, m := len(heightMap), len(heightMap[0])
    res := 0
    totalVisited := make(map[key] bool)
    
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            heap.Push(h, key{ i, j, heightMap[i][j] })
        }
    }
    
    for h.Len() != 0 {
        pop := heap.Pop(h).(key)
        stack := []key{pop}
        visited := make(map[key] bool)
        canFill := true
        min := 20000 // max value of heightMap[i][j] is 2 * 10^4, or 20000
        max := pop.height
        
        if totalVisited[pop] { continue }
        totalVisited[pop] = true
        visited[pop] = true
        
        helper := func(i, j, height int) bool {
            if i < 0 || j < 0 || i >= n || j >= m {
                return false
            }
            
            k := key{ i, j, heightMap[i][j] }
            
            if heightMap[i][j] <= height && !visited[ k ] {
                stack = append(stack, k)
                visited[k] = true
            } else if heightMap[i][j] > height && heightMap[i][j] < min {
                min = heightMap[i][j]
            }
            
            return true
        }
        
        for len(stack) != 0 {
            p := stack[len(stack) - 1]
            stack = stack[:len(stack) - 1]
            i, j, height := p.i, p.j, p.height
            
            if !helper(i + 1, j, height) ||
            !helper(i - 1, j, height) ||
            !helper(i, j + 1, height) || 
            !helper(i, j - 1, height) {
                canFill = false
                break
            }
        }
        
        if canFill {
            temp := min - max
            res += len(visited) * temp
            
            for a := range visited {
                totalVisited[a] = true
                heightMap[a.i][a.j] += temp
                heap.Push(h, key{ a.i, a.j, a.height + temp })
            }
        } else {
            for a := range visited {
                totalVisited[a] = true
            }
        }
    }
    
    return res
}

// work up layers


