type key struct {
    arr []int
    canNotDo bool
}

type key2 struct {
    a, b int
}

func countCompleteComponents(n int, edges [][]int) int {
    parents := make([]int, n)
	for i := 0; i < n; i++ {
		parents[i] = i
	}
	u := &unionFind{
		count:   n,
		parents: parents,
	}
    
    m := map[key2] bool {}
    
    for _, edge := range edges {
        u.Union(edge[0], edge[1])
        m[key2{edge[0], edge[1]}] = true
        m[key2{edge[1], edge[0]}] = true
    }
    
    res := u.count
    
    x := map[int] key {} // parent : children
    
    for i := 0; i < n; i++ {
        z := u.Find(i)
        
        if x[z].canNotDo {
            continue
        }
        
        canNotDo := true
        
        for _, a := range x[z].arr {
            if !m[key2{a, i}] {
                canNotDo = false
            }
        }
        
        if !canNotDo {
            res--
            l := key{x[z].arr, true}
            x[z] = l
        } else {
            l := key{append(x[z].arr, i), x[z].canNotDo}
            x[z] = l
        }
    }
    
    return res
}

type unionFind struct {
	count   int
	parents []int
}

func (this *unionFind) Find(x int) int {
	if x == this.parents[x] {
		return x
	}
	this.parents[x] = this.Find(this.parents[x])
	return this.parents[x]
}

func (this *unionFind) Union(x, y int) {
	rx, ry := this.Find(x), this.Find(y)
	if rx != ry {
		this.parents[rx] = ry
		this.count--
	}
}
