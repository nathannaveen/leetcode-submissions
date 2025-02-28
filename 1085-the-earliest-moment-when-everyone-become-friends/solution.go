type unionFind struct {
	parent []int
	n      int
}

func (this *unionFind) Find(x int) int {
	if x != this.parent[x] {
        this.parent[x] = this.Find(this.parent[x])
    	return this.parent[x]
	}

	return x
}

func (this *unionFind) Union(x int, y int) {
	temp1 := this.Find(x)
    temp2 := this.Find(y)
    
	if temp1 != temp2 {
        this.n--
	    this.parent[temp1] = temp2
	}
}

func earliestAcq(logs [][]int, N int) int {
    sort.SliceStable(logs, func(i, j int) bool { // Sort by their timestamp
        return logs[i][0] < logs[j][0]
    })
    
    u := &unionFind{ make([]int, N), N }
    
    for i := 0; i < N; i++ {
        u.parent[i] = i
    }
    
    for _, log := range logs {
        u.Union(log[1], log[2])
        
        if u.n == 1 {
            return log[0]
        }
    }
    
    return -1
}
