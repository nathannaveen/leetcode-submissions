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

func (this *unionFind) Union(x int, y int) bool {
	temp1 := this.Find(x)
	temp2 := this.Find(y)

	if temp1 != temp2 {
		this.n--
		this.parent[temp1] = temp2
        return true
    } else {
        return false // This signifies there is a redundent connection
    }
}

func findRedundantConnection(edges [][]int) []int {
    res := []int{}
    
    u := unionFind{ make([]int, len(edges)), len(edges) }
    
    for i := 0; i < len(edges); i++ {
        u.parent[i] = i
    }
    
    for _, edge := range edges {
        if !u.Union(edge[0] - 1, edge[1] - 1) { // checking for connections
            res = edge
        }
    }
    
    return res
}