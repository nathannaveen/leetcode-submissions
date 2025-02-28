type MaxHeap []float64

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(float64)) }

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Main Code Starts Bellow:

func halveArray(nums []int) int {
    h := &MaxHeap{}
    heap.Init(h)
    sum := float64(0)
    
    for _, num := range nums {
        heap.Push(h, float64(num))
        sum += float64(num)
    }
    
    k := float64(0)
    i := 0
    
    for sum - k > sum / 2 {
        pop := heap.Pop(h).(float64)
        k += pop / 2
        heap.Push(h, pop / 2)
        i++
    }
    
    return i
}