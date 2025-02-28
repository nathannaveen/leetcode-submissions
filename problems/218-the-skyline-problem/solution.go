type key struct {
    start bool // if it is a start true, if end false
    x int
    height int
    id int
}

type MinHeap []key

func (h MinHeap) Len() int                { return len(h) }
func (h MinHeap) Less(i, j int) bool { 
    if h[i].x == h[j].x {
        return h[i].height > h[j].height
    }
    
    return h[i].x < h[j].x 
}
func (h MinHeap) Swap(i, j int)           { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(value interface{}) { *h = append(*h, value.(key)) }
func (h *MinHeap) Pop() interface{} {
	min := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return min
}

func getSkyline(buildings [][]int) [][]int {
    h := &MinHeap{}
    heap.Init(h)
    ends := map[int] int {} // id : height
    
    newBuildings := [][]int{}
    
    for _, building := range buildings {
        done := false
        for i := 0; i < len(newBuildings); i++ {
            if building[0] <= newBuildings[i][1] && building[2] == newBuildings[i][2] {
                if building[1] > newBuildings[i][1] {
                    newBuildings[i][1] = building[1]
                }
                done = true
            }
        }
        
        if !done {
            newBuildings = append(newBuildings, []int{building[0], building[1], building[2]})
        }
    }
    
    for i, building := range newBuildings {
        heap.Push(h, key{ true, building[0], building[2], i}) // starting
        heap.Push(h, key{ false, building[1] + 1, building[2], i}) // ending
        ends[i] = building[2]
    }
    
    arr := []key{}
    
    m := make(map[int] int) // id : height
    
    for h.Len() > 0 {
        pop := heap.Pop(h).(key)
        
        if pop.start {
            m[pop.id] = pop.height
        } else {
            delete(m, pop.id)
        }
        
        max := 0

        for _, b := range m {
            if b > max {
                max = b
            }
        }

        arr = append(arr, key{pop.start, pop.x, max, pop.id})
    }
    
    for i := 0; i < len(arr); i++ {
        for j := 0; j < len(arr); j++ {
            if !arr[i].start && !arr[j].start && arr[i].x == arr[j].x && arr[j].height > arr[i].height {
                arr = append(arr[:j], arr[j + 1:]...)
                i -= 2
                j -= 2
            }
        }
    }
    
    res := [][]int{}
    
    for i := 0; i < len(arr); i++ {
        if arr[i].start {
            if arr[i].height > ends[arr[i].id] {
                continue
            }
            res = append(res, []int{arr[i].x, arr[i].height})
        } else {
            if arr[i].height > ends[arr[i].id] {
                continue
            }
            res = append(res, []int{arr[i].x - 1, arr[i].height})
        }
    }
    
    return res
}