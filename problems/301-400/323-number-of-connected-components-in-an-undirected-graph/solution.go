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

func (this *unionFind) Union(x int, y int) {
	rx, ry := this.Find(x), this.Find(y)
	if rx == ry {
		return
	}

	this.count--
	this.parents[rx] = ry
}

func countComponents(n int, edges [][]int) int {
    u := unionFind{ make([]int, n), n }

    for i := 0; i < n; i++ {
        u.parents[i] = i
    }
    
    for _, edge := range edges {
        u.Union(edge[0], edge[1])
    }
    return u.count
}