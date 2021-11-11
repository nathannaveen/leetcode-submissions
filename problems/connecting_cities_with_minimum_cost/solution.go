type unionFind struct {
	parents []int
	count   int
}

func (this *unionFind) Find(x int) int {
	if x == this.parents[x] {
		return x
	}

	this.parents[x] = this.Find(this.parents[x])
	return this.parents[x]
}

func (this *unionFind) Union(x int, y int) bool {
	rx, ry := this.Find(x), this.Find(y)
	if rx == ry {
		return false // redundent connection
	}

	this.count--
	this.parents[rx] = ry
    return true // made new connection
}

func minimumCost(n int, connections [][]int) int {
	parents := make([]int, n)
    res := 0
    
	for i := 0; i < n; i++ {
		parents[i] = i
	}
    
	unionFind := &unionFind{ parents, n }
    
    sort.Slice(connections, func(i, j int) bool {
        return connections[i][2] < connections[j][2]
    })
    
    for i := 0; i < len(connections); i++ {
        if unionFind.Union(connections[i][0] - 1, connections[i][1] - 1) {
            res += connections[i][2]
        }
    }
    
    if unionFind.count != 1 {
        return -1
    }
    
    return res
}