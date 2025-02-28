type key struct {
    x int
    y int
    dist float64
}

func kClosest(points [][]int, k int) [][]int {
    arr := []key{}
    res := [][]int{}
    
    for _, point := range points {
        x := point[0]
        y := point[1]
        
        arr = append(arr, key{ x, y, math.Sqrt(math.Pow(float64(x), 2) + math.Pow(float64(y), 2)) })
    }
    sort.Slice(arr, func(i, j int) bool { return arr[i].dist < arr[j].dist })
    
    for i := 0; i < k; i++ {
        res = append(res, []int{ arr[i].x, arr[i].y })
    }
    
    return res
}