type key struct {
    name string
    height int
}

func sortPeople(names []string, heights []int) []string {
    arr := []key{}
    
    for i := range names {
        arr = append(arr, key{ names[i], heights[i] })
    }
    
    sort.Slice(arr, func(i, j int) bool {
        return arr[i].height > arr[j].height
    })
    
    for i := range arr {
        names[i] = arr[i].name
    }
    
    return names
}