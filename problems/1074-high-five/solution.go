// type array struct {
//     arr []int
// }

func highFive(items [][]int) [][]int {
    m := make(map[int] []int)
    res := [][]int{}
    
    for _, i := range items {
        m[i[0]] = append(m[i[0]], i[1])
    }
    
    for a, b := range m {
        sort.Ints(b)
        sum := 0
        for i := int(math.Max(float64(len(b) - 5), float64(0))); i < len(b); i++ {
            sum += b[i]
        }
        
        res = append(res, []int{a, sum / int(math.Min(float64(len(b)), float64(5)))})
    }
    
    sort.Slice(res, func(i, j int) bool {
		return res[i][0] < res[j][0]
	})
    
    return res
}