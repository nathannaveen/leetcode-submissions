func countPairs(n int, edges [][]int) int64 {
    parents := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parents[i] = i
		size[i] = 1
	}
	u := &unionFind{
		count:   n,
		parents: parents,
		size: size,
	}

	res := int64((n - 1) * (n) / 2)

    for _, edge := range edges {
        x := u.Union(edge[0], edge[1])
		res -= int64(x)
	}

    return res
}

type unionFind struct {
	count   int
	parents []int
	size []int // each unions size
}

func (this *unionFind) Find(x int) int {
	if x == this.parents[x] {
		return x
	}
	this.parents[x] = this.Find(this.parents[x])
	return this.parents[x]
}

func (this *unionFind) Union(x, y int) int {
	rx, ry := this.Find(x), this.Find(y)
	if rx != ry {
		this.parents[rx] = ry
		x := this.size[rx] * this.size[ry]
		this.size[ry] += this.size[rx]
		this.count--
		return x
	}
	return 0
}
