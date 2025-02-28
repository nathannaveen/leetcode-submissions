type key struct {
    i int
    n int
}

func sortEvenOdd(nums []int) []int {
    even, odd := []int{}, []int{}
    
    for i, n := range nums {
        if i % 2 == 0 {
            even = append(even, n)
        } else {
            odd = append(odd, n)
        }
    }
    
    sort.Ints(even)
    
    sort.Slice(odd, func(i, j int) bool {
        return odd[i] > odd[j]
    })
    
    res := []int{}
    
    for i := 0; i < int(math.Min(float64(len(even)), float64(len(odd)))); i++ {
        res = append(res, even[i], odd[i])
    }
    
    if len(odd) != len(even) {
        res = append(res, even[len(even) - 1])
    }
    
    return res
}