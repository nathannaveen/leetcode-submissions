func mostFrequentIDs(nums []int, freq []int) []int64 {
    curValues := map[int] int64 {} // value : frequency
    h := &KeyHeap{}
    heap.Init(h)
    res := make([]int64, len(nums))

    for i := 0; i < len(nums); i++ {
        curValues[nums[i]] += int64(freq[i])
        heap.Push(h, key{value: nums[i], freq: curValues[nums[i]]})

        for curValues[(*h)[0].value] != (*h)[0].freq {
            heap.Pop(h)
        }

        res[i] = (*h)[0].freq
    }

    return res
}

type key struct {
    value int
    freq int64
}

type KeyHeap []key

func (h KeyHeap) Len() int           { return len(h) }
func (h KeyHeap) Less(i, j int) bool { return h[i].freq > h[j].freq }
func (h KeyHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *KeyHeap) Push(x any) {
	*h = append(*h, x.(key))
}

func (h *KeyHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
