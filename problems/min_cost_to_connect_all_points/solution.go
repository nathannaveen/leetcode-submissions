var res = 0

type key struct {
    one [2]int
    two [2]int
    distance int
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

func (this *unionFind) Union(x, y, val int) {
	temp1 := this.Find(x)
	temp2 := this.Find(y)

	if temp1 != temp2 {
        res += val
		this.n--
		this.parent[temp1] = temp2
	}
}

func minCostConnectPoints(points [][]int) int {
    u := unionFind{ []int{}, len(points) }
    m := make(map[[2]int] int)
    
    for i := range points {
        u.parent = append(u.parent, i)
        m[[2]int{ points[i][0], points[i][1] }] = i
    }
    
    arr := []key{}
    res = 0
	
    for i := 0; i < len(points); i++ {
        for j := i + 1; j < len(points); j++ {
            arr = append(arr, key{ [2]int{ points[i][0], points[i][1] },
                                  [2]int{ points[j][0], points[j][1] },
                                  abs(points[i][0] - points[j][0]) + abs(points[i][1] - points[j][1]) })
        }
    }
    
    sort.Slice(arr, func(i, j int) bool {
        return arr[i].distance < arr[j].distance
    })
    
    for i := range arr {
        u.Union(m[arr[i].one], m[arr[i].two], arr[i].distance)
    }
    
	return res
}

func abs(a int) int {
    return int(math.Abs(float64(a)))
}
