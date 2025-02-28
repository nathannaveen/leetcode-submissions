type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func largestInteger(num int) int {
    odd := &IntHeap{}
    even := &IntHeap{}
    heap.Init(odd)
    heap.Init(even)
    res := 0
    
    n := num
    
    for n > 0 {
        digit := n % 10
        
        if digit % 2 == 0 {
            heap.Push(even, digit)
        } else {
            heap.Push(odd, digit)
        }
        
        n /= 10
    }
    
    tmp := 1
    
    for num > 0 {
        digit := num % 10
        
        if digit % 2 == 0 {
            res += heap.Pop(even).(int) * tmp
        } else {
            res += heap.Pop(odd).(int) * tmp
        }
        
        tmp *= 10
        num /= 10
    }
    
    return res
}