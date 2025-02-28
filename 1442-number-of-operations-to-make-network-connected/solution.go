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

func (this *unionFind) Union(x int, y int) int {
    temp1 := this.Find(x)
    temp2 := this.Find(y)

    if temp1 != temp2 {
        this.n--
        this.parent[temp1] = temp2
        return 0
    } else {
        return 1 // This signifies there is a redundant connection
    }
}

func makeConnected(n int, connections [][]int) int {
    u := unionFind{ make([]int, n), n }
    res := 0
    
    for i := 0; i < n; i++ {
        u.parent[i] = i
    }
    
    for _, connection := range connections {
        res += u.Union(connection[0], connection[1])
    }
    
    if res >= u.n - 1 {
        return u.n - 1
    }
    return -1
}