type key struct {
    hours int
    minutes int
}

func findHighAccessEmployees(access_times [][]string) []string {
    res := []string{}
    m := map[string] *KeyHeap {}

    for _, access := range access_times {
        if _, ok := m[access[0]]; !ok {
            m[access[0]] = &KeyHeap{}
            heap.Init(m[access[0]])
        }
        hours, _ := strconv.Atoi(access[1][:2])
        minutes, _ := strconv.Atoi(access[1][2:])
        heap.Push(m[access[0]], key{hours, minutes})
    }

    for id, h := range m {
        if h.Len() < 3 {
            continue
        }

        first := heap.Pop(h).(key)
        second := heap.Pop(h).(key)

        for h.Len() > 0 {
            pop := heap.Pop(h).(key)

            if isInHourDiff(first.hours, first.minutes, pop.hours, pop.minutes) {
                res = append(res, id)
                break
            }

            first, second = second, pop
        }
    }

    return res
}

func isInHourDiff(h1, m1, h2, m2 int) bool {
    return h2 == h1 || (h2 == h1 + 1 && m2 < m1)
}


type KeyHeap []key

func (h KeyHeap) Len() int           { return len(h) }
func (h KeyHeap) Less(i, j int) bool { 
    if h[i].hours == h[j].hours {
        return h[i].minutes < h[j].minutes
    }
    return h[i].hours < h[j].hours
}
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