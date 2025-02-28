type key struct {
    index, cnt int
}

func maximumBeauty(nums []int, k int) int {
    indexes := map[int] int {} // number : index
    lineSweep := []key{}
    
    for _, num := range nums {
        if _, ok := indexes[num - k]; !ok { // if indexes doesn't contain (num - k)
            lineSweep = append(lineSweep, key{num - k, 0})
            indexes[num - k] = len(lineSweep) - 1
        }
        if _, ok := indexes[num + k + 1]; !ok { // if indexes doesn't contain (num + k + 1)
            lineSweep = append(lineSweep, key{ num + k + 1, 0 })
            indexes[num + k + 1] = len(lineSweep) - 1
        }
        
        lineSweep[indexes[num - k]].cnt++
        lineSweep[indexes[num + k + 1]].cnt--
    }
    
    sort.Slice(lineSweep, func(i, j int) bool {
        // sort lineSweep by each index so that we can actualy implement the line sweep
        return lineSweep[i].index < lineSweep[j].index
    })
    
    sum := 0
    max := 0
    
    for _, a := range lineSweep {
        sum += a.cnt
        if sum > max {
            max = sum
        }
    }
    
    return max
}