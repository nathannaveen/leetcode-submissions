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

func findCircleNum(isConnected [][]int) int {
    u := unionFind{ make([]int, len(isConnected)), len(isConnected) }
    
    for i := range isConnected {
        u.parents[i] = i
    }
    
    for i := range isConnected {
        for j := range isConnected[i] {
            if isConnected[i][j] == 1 {
                u.Union(i, j)
            }
        }
    }
    return u.count
}