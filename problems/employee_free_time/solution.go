func employeeFreeTime(schedule [][]*Interval) []*Interval {
    arr := []*Interval{}
    res := []*Interval{}
    
    for _, ints := range schedule {
        for _, i := range ints {
            arr = append(arr, i)
        }
    }
    
    sort.Slice(arr, func(i, j int) bool {
        return arr[i].Start <= arr[j].Start
    })
    
    for i := 0; i < len(arr) - 1; i++ {
        if arr[i].End < arr[i + 1].Start {
            continue
        }
        if arr[i].End >= arr[i + 1].End {
            arr = append(arr[:i + 1], arr[i + 2:]...)
        } else {
            arr[i].End = arr[i + 1].End
            arr = append(arr[:i + 1], arr[i + 2:]...)
        }
        i--        
    }
    
    for i := 0; i < len(arr) - 1; i++ {
        res = append(res, &Interval{ arr[i].End, arr[i + 1].Start })
    }
    
    return res
}