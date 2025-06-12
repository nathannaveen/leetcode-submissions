func minimumIndex(nums []int) int {
    indexes := map[int] []int{}
    m := map[int] int {}
    
    maxVal := 0
    
    for i, num := range nums {
        indexes[num] = append(indexes[num], i)
        m[num]++
        if m[num] > m[maxVal] {
            maxVal = num
        }
    }
    
    for i, index := range indexes[maxVal] {
        // i : number of the value
        if (i + 1) * 2 > index + 1 && (len(indexes[maxVal]) - i - 1) * 2 > len(nums) - index - 1 {
            return index
        }
    }
    
    return -1
}
