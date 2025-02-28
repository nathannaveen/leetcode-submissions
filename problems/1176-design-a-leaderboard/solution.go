type key struct {
    id int
    score int
}

type Leaderboard struct {
    h *KeyHeap
    value map[int] int // id : score
}


func Constructor() Leaderboard {
    h := &KeyHeap{}
    heap.Init(h)
    return Leaderboard{ h, map[int] int {} }
}


func (this *Leaderboard) AddScore(playerId int, score int)  {
    this.value[playerId] += score
    heap.Push(this.h, key{playerId, this.value[playerId]})
}


func (this *Leaderboard) Top(K int) int {
    sum := 0
    addBack := make([]key, K)

    for K > 0 {
        pop := heap.Pop(this.h).(key)
        if this.value[pop.id] == pop.score {
            sum += pop.score
            addBack[len(addBack) - K] = pop
            K--
        }
    }

    for _, push := range addBack {
        heap.Push(this.h, push)
    }

    return sum
}


func (this *Leaderboard) Reset(playerId int)  {
    this.value[playerId] = 0
}

type KeyHeap []key

func (h KeyHeap) Len() int           { return len(h) }
func (h KeyHeap) Less(i, j int) bool { return h[i].score > h[j].score }
func (h KeyHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *KeyHeap) Push(x interface{}) {*h = append(*h, x.(key))}

func (h *KeyHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

/**
 * Your Leaderboard object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddScore(playerId,score);
 * param_2 := obj.Top(K);
 * obj.Reset(playerId);
 */