type key struct {
    i int
    j int
}

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

func numIslands2(m int, n int, positions [][]int) []int {
    u := unionFind{ []int{}, 0 }
    res := make([]int, len(positions))
    m2 := make(map[key] int)
    
    for i := range positions {
        u.parent = append(u.parent, i)
    }
    
    for i, position := range positions {
        if m2[ key{ position[0], position[1] } ] == 0 { // if we have not already made a island on the position
            if m2[ key{ position[0] - 1, position[1] } ] != 0 {
                u.Union(i, m2[ key{ position[0] - 1, position[1] } ] - 1)
            } 
            if m2[ key{ position[0] + 1, position[1] } ] != 0 {
                u.Union(i, m2[ key{ position[0] + 1, position[1] } ] - 1)
            } 
            if m2[ key{ position[0], position[1] - 1 } ] != 0 {
                u.Union(i, m2[ key{ position[0], position[1] - 1 } ] - 1)
            } 
            if m2[ key{ position[0], position[1] + 1 } ] != 0 {
                u.Union(i, m2[ key{ position[0], position[1] + 1 } ] - 1)
            }
            
            m2[ key{ position[0], position[1] } ] = i + 1 // there is an island at position i
            
            u.n++ // add new island to count
        }
        
        res[i] = u.n
        
    }
    return res
}